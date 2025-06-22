package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttpEcho(t *testing.T) {
	http.HandleFunc("POST /echo", EchoHandler)
	rr := httptest.NewRecorder()
	body := strings.NewReader("hello")
	req, err := http.NewRequest(http.MethodPost, "/echo", body)
	if err != nil {
		t.Fatal(err)
	}
	http.DefaultServeMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Response code is not ok: %v", rr.Code)
	}
	if expected, actual := "hello", rr.Body.String(); expected != actual {
		t.Errorf("Wrong request body. Expected: %s, actual: %s", expected, actual)
	}
}
