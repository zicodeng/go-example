package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// To test our HelloHandler, we need actually send an HTTP request with query
	// to the path that will trigger the HelloHandler.
	req := httptest.NewRequest("GET", "/hello", nil)

	// Define query (key-value pair).
	q := req.URL.Query()
	name := "zico"
	q.Add("name", name)
	// Set query in request URL.
	req.URL.RawQuery = q.Encode()

	// NewRecorder is an implementation of http.ResponseWriter.
	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(HelloHandler)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned the wrong status code: got %d want %d", status, http.StatusOK)
	}

	expectedOutput := fmt.Sprintf("Hello, %s!", name)
	if recorder.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body: got %s want %s", recorder.Body.String(), expectedOutput)
	}
}
