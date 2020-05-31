package testing

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// spins up the server and tests the routing too
func TestServer(t *testing.T) {
	s := httptest.NewServer(NewServer())
	defer s.Close()

	// Create request to the server
	req, err := http.NewRequest("GET", s.URL+"/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("expected no error, got: %s", err.Error())
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("router unable to find route")
	}
}

// This test will test the handler function directly
// Achieved by creating a handler func on the method we want to test
// then serving a request to that handler and checking against what was recorded
func TestHandler(t *testing.T) {
	// Create request to the server
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set up recorder and handler
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(helloWorld)

	handler.ServeHTTP(res, req)

	// Check the status code is what we expect.
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `Hello, world!`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}
