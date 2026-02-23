package server

import "net/http"

type Server struct {
	Service Service
	Addr string
	Mux *http.ServeMux
}

func NewServer(addr string, s Service) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", PingHandler)
	return &Server{
		Service: s,
		Addr: addr,
		Mux: mux,
	}
}
