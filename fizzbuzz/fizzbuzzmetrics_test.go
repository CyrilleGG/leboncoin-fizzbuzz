package fizzbuzz

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)




//		Function to test hashParams()
func TestHashParams (t *testing.T) {

	//		Given
	want := ParamsTracker {
		IpAddress:    "111.111.111.111",
		FirstInt:     3,
		SecondInt:    5,
		Limit:        100,
		FirstString:  "fizz",
		SecondString: "buzz",
		ParamsHash:   "56e6c4a0ed1e2c6cd9f330aa576235be",
	}

	//		When
	got := ParamsTracker {
		IpAddress:    "111.111.111.111",
		FirstInt:     3,
		SecondInt:    5,
		Limit:        100,
		FirstString:  "fizz",
		SecondString: "buzz",
		ParamsHash:   "",
	}
	err := got.hashParams()

	//		Then
	if err != nil {
		t.Errorf("hashParams() failed: %v", err)
	}

	if want.ParamsHash != got.ParamsHash {
		t.Errorf("hashParams() failed. Got %v instead of %v", got.ParamsHash, want.ParamsHash)
	}
}




//		Function to test GetMetrics()
func TestGetMetrics(t *testing.T) {

	//		Creating mock request with custom params
	//		and checking for error
	request, err := http.NewRequest("GET", "/fizzbuzz-metrics", nil)
	if err != nil {
		t.Errorf("GetMetrics() failed because of bad request: %v", err)
	}

	//		Creating a HTTP recorder for the test
	recorder := httptest.NewRecorder()

	//		Hitting endpoint
	router := httprouter.New()
	router.GET("/fizzbuzz-metrics", GetMetrics)
	router.ServeHTTP(recorder, request)

	//		Check status code
	if recorder.Code != http.StatusOK {
		t.Errorf("GetMetrics() failed because of status code: got %v instead of %v", recorder.Code, http.StatusOK)
	}

	//		Check body
	got := recorder.Body.String()
	if got == "" {
		t.Error("GetMetrics() failed because of empty body.")
	}
}