package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttpEcho(t *testing.T) {
	// Arrange
	http.HandleFunc("POST /echo", EchoHandler)
	rr := httptest.NewRecorder()
	data := "hello"
	reqBody := strings.NewReader(data)
	req, err := http.NewRequest(http.MethodPost, "/echo", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	// Act
	http.DefaultServeMux.ServeHTTP(rr, req)

	// Assert
	if rr.Body.String() != data {
		t.Errorf("Wrong response body. Expected: %s, got: %s", data, rr.Body.String())
	}
	if rr.Code != http.StatusOK {
		t.Errorf("Wrong status code. Expected: %d, got: %d", http.StatusOK, rr.Code)
	}

	// Arrange
	rr = httptest.NewRecorder()
	req, err = http.NewRequest(http.MethodGet, "/echo", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Act
	http.DefaultServeMux.ServeHTTP(rr, req)

	// Assert
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Wrong status code. Expected: %d, got: %d", http.StatusMethodNotAllowed, rr.Code)
	}
}
