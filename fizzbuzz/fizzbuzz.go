package fizzbuzz

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/julienschmidt/httprouter"

	"github.com/cyrillegg/leboncoin-fizzbuzz/server"
)




//		Defining needed custom variables for fizzbuzz
//		function
var limit 		int		= 100
var int1 		int		= 3
var int2 		int		= 5
var str1 		string	= "fizz"
var str2 		string	= "buzz"

//		Defining payload struct that will contain data
//		in GO data format before converting in JSON
type payload struct {
	Results 	[]string 	`json:"fizzbuzz_results"`
}

//		Defining err variable for handling errors
var err	error




//		Fizzbuzz function with custom arguments
func fizzbuzz (i1 int, i2 int, l int, s1 string, s2 string) []string {

	//		Defining temporary variables
	var count 		int = 1
	var results 	[]string

	//		Looping from 1 to limit
	for count <= l {

		//		Conditions checking potential multiples
		//		and replacing them by appropriate strings
		if count%i1 == 0 && count%i2 == 0 {
			results = append(results, s1+s2)
		} else if count%i1 == 0 {
			results = append(results, s1)
		} else if count%i2 == 0 {
			results = append(results, s2)
		} else {
			results = append(results, strconv.Itoa(count))
		}

		//		Increment count to continue loop
		count++
	}

	//		Returning final results
	return results
}




//		Route's main function
func Res(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	//		Defining response's header's configuration
	var header = writer.Header()
	header.Set("Content-Type", "application/json")

	//		Setting prefix and flags for logger
	log.SetPrefix("fizzbuzz.go: ")
	log.SetFlags(0)

	//		Getting params from the HTTP query
	q := request.URL.Query()

	//		Assigning params' values to custom variables with
	//		error handling
	if q.Get("first_int") != "" {
		int1, err = strconv.Atoi(q.Get("first_int"))
		if server.CheckError(writer, err, http.StatusBadRequest) != nil {
			return
		}
	}

	if q.Get("second_int") != "" {
		int2, err = strconv.Atoi(q.Get("second_int"))
		if server.CheckError(writer, err, http.StatusBadRequest) != nil {
			return
		}
	}

	//		Making sure that int1 > 0 and int2 > 0
	if int1 == 0 || int2 == 0 {
		log.Println("int1 and int2 cannot be 0")
		http.Error(writer, "Please, choose a value above 0 for 'first_int' and 'second_int'", http.StatusBadRequest)
		return
	}

	if q.Get("limit") != "" {
		limit, err = strconv.Atoi(q.Get("limit"))
		if server.CheckError(writer, err, http.StatusBadRequest) != nil {
			return
		}
	}

	if q.Get("first_string") != "" {
		str1 = q.Get("first_string")
	}

	if q.Get("second_string") != "" {
		str2 = q.Get("second_string")
	}

	//		Inserting metrics regarding params into DB
	//		with a goroutine
	var tracker = ParamsTracker {
		IpAddress:    request.RemoteAddr,
		FirstInt:     int1,
		SecondInt:    int2,
		Limit:        limit,
		FirstString:  str1,
		SecondString: str2,
		ParamsHash:   "",
	}
	var wg sync.WaitGroup
	wg.Add(1)
	err = tracker.Insert(&wg)
	if server.CheckError(writer, err, http.StatusInternalServerError) != nil {
		return
	}

	//		Declaring response's body that will contain
	//		the payload
	var body = payload {
		Results: fizzbuzz(int1, int2, limit, str1, str2),
	}

	//		Converting body into json
	dataset, err := json.Marshal(body)
	if server.CheckError(writer, err, http.StatusInternalServerError) != nil {
		return
	}

	//		Creating response
	var res = server.NewResponse(http.StatusOK, "OK", dataset)

	//		Waiting for WaitGroup to be done
	wg.Wait()

	//		Writing header and body of HTTP response
	header.Add("Message", res.Message)
	writer.WriteHeader(res.Status)
	_, err = writer.Write(res.Data)
	if server.CheckError(writer, err, http.StatusInternalServerError) != nil {
		return
	}
}