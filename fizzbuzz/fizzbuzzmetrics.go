package fizzbuzz

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/julienschmidt/httprouter"

	"github.com/cyrillegg/leboncoin-fizzbuzz/database"
	"github.com/cyrillegg/leboncoin-fizzbuzz/server"
)




//		Defining ParamsTracker struct
type ParamsTracker struct {
	IpAddress		string
	FirstInt		int
	SecondInt		int
	Limit			int
	FirstString		string
	SecondString	string
	ParamsHash		string
}

//		Defining Metrics struct
type Metrics struct {
	FirstInt		int		`json:"first_int"`
	SecondInt		int		`json:"second_int"`
	Limit			int		`json:"limit"`
	FirstString		string	`json:"first_string"`
	SecondString	string	`json:"second_string"`
	ParamsHash		string	`json:"params_hash"`
	Count			int		`json:"hits"`
}




//		Function to hash data
func (m *ParamsTracker) hashParams() error {

	//		Converting integers into strings
	var int1 = strconv.Itoa(m.FirstInt)
	var int2 = strconv.Itoa(m.SecondInt)
	var limit = strconv.Itoa(m.Limit)

	//		Preparing params string before hashing
	var params = "first_int: " + int1 + "second_int: " + int2 + "limit: " + limit + "first_string: " + m.FirstString + "second_string: " + m.SecondString

	//		Hashing and converting hash to string
	var hParams = md5.New()
	hParams.Write([]byte(params))
	var hexaParams = hParams.Sum(nil)
	var hashParamsStr = hex.EncodeToString(hexaParams)

	//		Assigning result
	m.ParamsHash = hashParamsStr

	return nil
}




//		Function to record used params into DB
func (m *ParamsTracker) Insert(waitgroup *sync.WaitGroup) error {
	err := m.hashParams()
	if err != nil {
		return err
	}

	//		Writing SQL query with custom params
	var query = `
INSERT INTO fizzbuzz_queries ("id", "ip_address", "time", "first_int", "second_int", "limit", "first_string", "second_string", "params_hash")
	VALUES (DEFAULT, $1, DEFAULT, $2, $3, $4, $5, $6, $7)
`

	//		Executing SQL query with params
	_, err = database.Open().Exec(query, m.IpAddress, m.FirstInt, m.SecondInt, m.Limit, m.FirstString, m.SecondString, m.ParamsHash)
	if err != nil {
		return err
	}

	//		Indicating that the goroutine is
	//		done and end function
	waitgroup.Done()
	return nil
}




//		Function to get metrics recorded on DB
func GetMetrics(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	//		Defining response's header's configuration
	var header = writer.Header()
	header.Set("Content-Type", "application/json")

	//		Setting prefix and flags for logger
	log.SetPrefix("fizzbuzzmetrics.go: ")
	log.SetFlags(0)

	//		Declaring response's body that will store
	//		the metrics
	var body = Metrics {
		FirstInt:     0,
		SecondInt:    0,
		Limit:        0,
		FirstString:  "",
		SecondString: "",
		ParamsHash:   "",
		Count:        0,
	}

	//		Writing SQL query
	query := `
SELECT first_int, second_int, "limit", first_string, second_string, params_hash, COUNT(params_hash) AS hits 
FROM fizzbuzz_queries  GROUP BY first_int, second_int, "limit", first_string, second_string, params_hash 
HAVING COUNT(params_hash)=( 
	SELECT MAX(params_cnt) 
	FROM ( 
		SELECT params_hash, COUNT(params_hash) AS params_cnt 
		FROM fizzbuzz_queries 
		GROUP BY params_hash
	) AS foo
);
`

	// 		Preparing query
	stmt, err := database.Open().Prepare(query)
	if server.CheckError(writer, err, http.StatusInternalServerError) != nil {
		return
	}
	defer stmt.Close()

	// 		Starting query
	rows, err := stmt.Query()
	if server.CheckError(writer, err, http.StatusInternalServerError) != nil {
		return
	}
	defer rows.Close()

	// 		Stocking data from SQL result into body
	for rows.Next() {
		err = rows.Scan(
			&body.FirstInt,
			&body.SecondInt,
			&body.Limit,
			&body.FirstString,
			&body.SecondString,
			&body.ParamsHash,
			&body.Count,
		)
		if server.CheckError(writer, err, http.StatusInternalServerError) != nil {
			return
		}
	}
	err = rows.Err()
	if server.CheckError(writer, err, http.StatusInternalServerError) != nil {
		return
	}

	//		Converting body into json
	dataset, err := json.Marshal(body)
	if server.CheckError(writer, err, http.StatusInternalServerError) != nil {
		return
	}

	//		Creating response
	var res = server.NewResponse(http.StatusOK, "OK", dataset)

	//		Writing header and body of HTTP response
	header.Add("Message", res.Message)
	writer.WriteHeader(res.Status)
	_, err = writer.Write(res.Data)
	if server.CheckError(writer, err, http.StatusInternalServerError) != nil {
		return
	}
}