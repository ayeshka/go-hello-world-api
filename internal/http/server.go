package http

import (
	"context"
	"log"
	"net/http"
	"time"
)

type timeOuts struct {
	read  int
	write int
	idle  int
}

type Server struct {
	httpServer *http.Server
	timeOuts   timeOuts
}

// NewServer creates a new Server instance with specified timeouts
func NewServer(readTimeOut, writeTimeout, idleTimeout int) *Server {
	return &Server{
		httpServer: &http.Server{},
		timeOuts: timeOuts{
			read:  readTimeOut,
			write: writeTimeout,
			idle:  idleTimeout,
		},
	}
}

// RegisterHandler registers an HTTP handler for a given pattern
func (s *Server) RegisterHandler(pattern string, handler http.Handler) {
	http.Handle(pattern, handler)
}

// Start starts the HTTP server on the specified port
func (s *Server) Start(port string) {
	// Start the HTTP server logic here
	s.httpServer = &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  time.Duration(s.timeOuts.read) * time.Second,
		WriteTimeout: time.Duration(s.timeOuts.write) * time.Second,
		IdleTimeout:  time.Duration(s.timeOuts.idle) * time.Second,
	}

	log.Println("Server starting on port 8080...")
	log.Println("Press Ctrl+C to gracefully shutdown the server")

	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", port, err)
	}
}

// Stop gracefully shuts down the server without interrupting any active connections
func (s *Server) Stop() {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
		return
	}
}
