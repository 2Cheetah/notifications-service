package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("POST /echo", EchoHandler)
	http.ListenAndServe(":8080", nil)
	slog.Info("Hello world")
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(body))
}
