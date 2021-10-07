package fizzbuzz

import (
	"encoding/json"
	"log"
	"strconv"

	"net/http"
	"github.com/julienschmidt/httprouter"

	//"../database"
	"../server"
)



//		Defining needed custom variables for fizzbuzz
//		function
var limit 		int
var int1 		int
var int2 		int
var str1 		string
var str2 		string

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
func Res(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	//		Defining response's header's configuration
	var header = w.Header()
	header.Set("Content-Type", "application/json")

	//		Getting params from the HTTP query
	q := r.URL.Query()

	//		Assigning params' values
	int1, err = strconv.Atoi(q.Get("first_int"))

	int2, err = strconv.Atoi(q.Get("second_int"))

	limit, err = strconv.Atoi(q.Get("limit"))

	str1 = q.Get("first_string")

	str2 = q.Get("second_string")

	//		Declaring response's body that will contain
	//		the payload
	var body = payload{
		Results: fizzbuzz(int1, int2, limit, str1, str2),
	}

	//		Converting body into json
	dataset, err := json.Marshal(body)
	if err != nil {
		log.Fatal("fizzbuzz.go line 68:", err)
		return
	}

	//		Creating response
	var res = server.NewResponse(200, "OK", dataset)

	//		Writing header and body of HTTP response
	w.WriteHeader(res.Status)
	_, err = w.Write(res.Data)
	if err != nil {
		log.Fatal("fizzbuzz.go line 81:", err)
		return
	}
}