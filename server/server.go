package server

import (
	"database/sql"
	"fmt"

	"github.com/julienschmidt/httprouter" // package for router
)

//		Defining server struct and its dependencies
type Server struct {
	Database *sql.DB
	Router   *httprouter.Router
}

//		Function used to create server's config with DB
func NewServer(database *sql.DB, router *httprouter.Router) (*Server, error){
	if database == nil {
		return nil, fmt.Errorf("DB config is not specified")
	} else if router == nil {
		return nil, fmt.Errorf("router is not specified")
	}
	s := &Server{
		Database: database,
		Router:   router,
	}

	return s, nil
}