package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	/*
		Here, we create a new HTTP request.
		This is the request that is going to be passad to our handler.
		The first argument is the method, the second is the route and the third,
		the request body.
	*/
	req, err := http.NewRequest("GET", "", nil)

	// In case there is an error creating the request, we fail and stop the test
	if err != nil {
		t.Fatal(err)
	}

	/*
		We use Go's httptest library to create an http recorder. This recorder
		will act as the target of our http request
		(its like a mini browser, which will accept the result of the request)
	*/
	recorder := httptest.NewRecorder()

	/*
		Create an HTTP handler from our handler function. "handler" is the handler
		function defined in our main.go file that we want to test
	*/

	hf := http.HandlerFunc(handler)

	/*
		Serve the HTTP request to our recorder. This is the line that actually
		executes the handler that we want to test.
	*/
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: expected %v actual %v", http.StatusOK, status)
	}

	// Check if the response body is what we expect
	expected := "Hello, Sir!"
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: expected %v got %v", expected, actual)
	}
}
