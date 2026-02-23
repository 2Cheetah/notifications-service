package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Pong int `json:"pong"`
	}{
		Pong: 1,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error("couldn't encode response to json", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
