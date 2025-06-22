package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("/", EchoHandler)
	http.ListenAndServe(":8080", nil)
	slog.Info("Hello world")
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	fmt.Fprint(w, string(body))
}
