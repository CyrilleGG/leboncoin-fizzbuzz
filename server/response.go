package server

import (
	"log"
	"net/http"
)

// 		Defining HTTPResponse struct for HTTP responses
type HTTPResponse struct {
	Status 	int 		`json:"statusCode"`
	Message string 		`json:"message"`
	Data 	[]byte		`json:"data"`
}

//		Function to create HTTP response
func NewResponse (status int, message string, data []byte) HTTPResponse {
	r := HTTPResponse {
		Status:  status,
		Message: message,
		Data:    data,
	}
	return r
}

//		Function that checks for any error from code and
//		sends a HTTP 500 response if so
func CheckError (writer http.ResponseWriter, e error, status int) interface{}{
	if e != nil {
		log.Println(e)
		http.Error(writer, e.Error(), status)
		return e
	}
	return nil
}
