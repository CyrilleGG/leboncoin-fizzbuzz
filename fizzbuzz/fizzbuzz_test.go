package fizzbuzz

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
)




//		Function to test fizzbuzz() with default arguments
func TestFizzbuzzDefaultArgs (t *testing.T) {

	//		Given
	want := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz", "fizz", "22", "23", "fizz", "buzz", "26", "fizz", "28", "29", "fizzbuzz", "31", "32", "fizz", "34", "buzz", "fizz", "37", "38", "fizz", "buzz", "41", "fizz", "43", "44", "fizzbuzz", "46", "47", "fizz", "49", "buzz", "fizz", "52", "53", "fizz", "buzz", "56", "fizz", "58", "59", "fizzbuzz", "61", "62", "fizz", "64", "buzz", "fizz", "67", "68", "fizz", "buzz", "71", "fizz", "73", "74", "fizzbuzz", "76", "77", "fizz", "79", "buzz", "fizz", "82", "83", "fizz", "buzz", "86", "fizz", "88", "89", "fizzbuzz", "91", "92", "fizz", "94", "buzz", "fizz", "97", "98", "fizz", "buzz"}

	//		When
	got := fizzbuzz(int1, int2, limit, str1, str2)

	//		Then
	for i, val := range got {
		if val != want[i] {
			t.Errorf("fizzbuzz() failed with default arguments. Got %#q instead of %#q", got, want)
		}
	}
}




//		Function to test Res() with no params
func TestResDefaultParams(t *testing.T) {

	//		Creating mock request with no params
	//		and checking for error
	request, err := http.NewRequest("GET", "/fizzbuzz", nil)
	if err != nil {
		t.Errorf("Res() failed because of bad request: %v", err)
	}

	//		Creating a HTTP recorder for the test
	recorder := httptest.NewRecorder()

	//		Hitting endpoint
	router := httprouter.New()
	router.GET("/fizzbuzz", Res)
	router.ServeHTTP(recorder, request)

	//		Check status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Res() failed because of status code: got %v instead of %v", recorder.Code, http.StatusOK)
	}

	//		Check body
	want := `{"fizzbuzz_results":["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz","16","17","fizz","19","buzz","fizz","22","23","fizz","buzz","26","fizz","28","29","fizzbuzz","31","32","fizz","34","buzz","fizz","37","38","fizz","buzz","41","fizz","43","44","fizzbuzz","46","47","fizz","49","buzz","fizz","52","53","fizz","buzz","56","fizz","58","59","fizzbuzz","61","62","fizz","64","buzz","fizz","67","68","fizz","buzz","71","fizz","73","74","fizzbuzz","76","77","fizz","79","buzz","fizz","82","83","fizz","buzz","86","fizz","88","89","fizzbuzz","91","92","fizz","94","buzz","fizz","97","98","fizz","buzz"]}`
	got := recorder.Body.String()
	if got != want {
		t.Errorf("Res() failed because of wrong body: got\n %v\n instead of\n %v", got, want)
	}
}




//		Function to test fizzbuzz() with custom arguments
func TestFizzbuzzCustomArgs (t *testing.T) {

	//		Given
	int1 = 3
	int2 = 8
	limit = 30
	str1 = "foo"
	str2 = "bar"
	want := []string{"1", "2", "foo", "4", "5", "foo", "7", "bar", "foo", "10", "11", "foo", "13", "14", "foo", "bar", "17", "foo", "19", "20", "foo", "22", "23", "foobar", "25", "26", "foo", "28", "29", "foo"}

	//		When
	got := fizzbuzz(int1, int2, limit, str1, str2)

	//		Then
	for i, val := range got {
		if val != want[i] {
			t.Errorf("fizzbuzz() failed with default values. Got %#q instead of %#q", got, want)
		}
	}
}




//		Function to test Res() with no params
func TestResCustomParams(t *testing.T) {

	//		Creating mock request with custom params
	//		and checking for error
	request, err := http.NewRequest("GET", "/fizzbuzz", nil)
	if err != nil {
		t.Errorf("Res() failed because of bad request: %v", err)
	}
	query := request.URL.Query()
	query.Add("first_int", "3")
	query.Add("second_int", "8")
	query.Add("limit", "30")
	query.Add("first_string", "foo")
	query.Add("second_string", "bar")

	//		Creating a HTTP recorder for the test
	recorder := httptest.NewRecorder()

	//		Hitting endpoint
	router := httprouter.New()
	router.GET("/fizzbuzz", Res)
	router.ServeHTTP(recorder, request)

	//		Check status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Res() failed because of status code: got %v instead of %v", recorder.Code, http.StatusOK)
	}

	//		Check body
	want := `{"fizzbuzz_results":["1","2","foo","4","5","foo","7","bar","foo","10","11","foo","13","14","foo","bar","17","foo","19","20","foo","22","23","foobar","25","26","foo","28","29","foo"]}`
	got := recorder.Body.String()
	if got != want {
		t.Errorf("Res() failed because of wrong body: got\n %v\n instead of\n %v", got, want)
	}
}
