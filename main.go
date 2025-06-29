package main

import (
	"fmt"
	"io"
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
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		slog.Error("failed to read request body", "error", err)
		return
	}
	fmt.Fprint(w, string(body))
}
