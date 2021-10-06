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



//		Defining needed variables for fizzbuzz
var count 		int = 1
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



//		Fizzbuzz function
func FizzBuzz(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	//		Defining response's header's configuration
	var header = w.Header()
	header.Set("Content-Type", "application/json")

	//		Declaring response's body that will contain
	//		the payload
	var body = payload{
		Results: 	[]string{},
	}

	limit = 30
	int1 = 3
	int2 = 5
	str1 = "fizz"
	str2 = "buzz"

	for count <= limit {
		if count%int1 == 0 && count%int2 == 0 {
			body.Results = append(body.Results, str1+str2)
		} else if count%int1 == 0 {
			body.Results = append(body.Results, str1)
		} else if count%int2 == 0 {
			body.Results = append(body.Results, str2)
		} else {
			body.Results = append(body.Results, strconv.Itoa(count))
		}

		count++
	}

	//		Converting body into json
	dataset, err := json.Marshal(body)
	if err != nil {
		log.Fatal("fizzbuzz.go line 68:", err)
		return
	}

	//		Creating response
	var res = server.HTTPResponse{
		Status: 	200,
		Message: 	"OK",
		Data: 		dataset,
	}

	//		Writing header and body of HTTP response
	w.WriteHeader(res.Status)
	_, err = w.Write(res.Data)
	if err != nil {
		log.Fatal("fizzbuzz.go line 81:", err)
		return
	}
}