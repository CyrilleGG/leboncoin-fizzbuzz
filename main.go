package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter" // package for router

	"github.com/cyrillegg/leboncoin-fizzbuzz/database"
	"github.com/cyrillegg/leboncoin-fizzbuzz/server"
	"github.com/cyrillegg/leboncoin-fizzbuzz/server/routes"
)

func main() {

	//		Declaring router with httprouter
	var router = httprouter.New()

	//		Creating server
	//		and check if error
	server, err := server.NewServer(database.Open(), router)
	if err != nil {
		log.Fatal(err)
		return
	}

	//		Launching API and its routes
	routes.Routes(server)
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":" + port, server.Router))
}
