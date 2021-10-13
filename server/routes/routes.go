package routes

import (
	"github.com/cyrillegg/leboncoin-fizzbuzz/fizzbuzz"
	"github.com/cyrillegg/leboncoin-fizzbuzz/server"
)

//		Listing API's routes
func Routes(s *server.Server) {
	s.Router.GET("/fizzbuzz", fizzbuzz.Res)
	s.Router.GET("/fizzbuzz-metrics", fizzbuzz.GetMetrics)
}