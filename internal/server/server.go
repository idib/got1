package server

import (
	"log"
	"net/http"
)

// Server represents the HTTP server
type Server struct {
	addr string
}

// NewServer creates a new server instance
func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

// Start initializes and starts the HTTP server
func (s *Server) Start() error {
	mux := http.NewServeMux()

	// Register the ping handler
	mux.HandleFunc("/ping", s.handlePing)

	log.Printf("Starting server on %s", s.addr)
	return http.ListenAndServe(s.addr, mux)
}

// handlePing handles the /ping endpoint
func (s *Server) handlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}