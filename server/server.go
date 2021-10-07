package server

import (
	//"database/sql"
	"fmt"

	"github.com/julienschmidt/httprouter"
)

//		Defining server struct and its dependencies
type Server struct {
	// Database *sql.DB
	Router   *httprouter.Router
}

// 		Defining HTTPResponse struct for HTTP responses
type HTTPResponse struct {
	Status 	int 		`json:"statusCode"`
	Message string 		`json:"message"`
	Data 	[]byte		`json:"data"`
}

//		Function to create HTTP response
func NewResponse (s int, m string, d []byte) HTTPResponse {
	r := HTTPResponse {
		Status:  s,
		Message: m,
		Data:    d,
	}
	return r
}

//		Function to create server's config
func NewServer(rt *httprouter.Router) (*Server, error){
	if rt == nil {
		return nil, fmt.Errorf("router is not specified")
	}
	s := &Server {
		Router:   rt,
	}

	return s, nil
}

//		Function used to create server's config with DB
//func NewServer(d *sql.DB, rt *httprouter.Router) (*Server, error){
//	if d == nil {
//		return nil, fmt.Errorf("DB config is not specified")
//	} else if rt == nil {
//		return nil, fmt.Errorf("router is not specified")
//	}
//	s := &Server{
//		Database: d,
//		Router:   rt,
//	}
//
//	return s, nil
//}