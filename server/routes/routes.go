package routes

import (
	".."

	"../../fizzbuzz"
)

//		Listing API's routes
func Routes(s *server.Server) {
	s.Router.GET("/fizzbuzz", fizzbuzz.FizzBuzz)
	//s.Router.GET("metrics", pathToFunction)
}