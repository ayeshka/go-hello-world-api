package main

import (
	"go-hello-world-api/internal/http"
	"go-hello-world-api/internal/http/handler"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Channel to listen for interrupt signals
	quit := make(chan os.Signal, 1)
	// Register the channel to receive specific signals
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Initialize and start the HTTP server
	// Here, we set read, write, and idle timeouts to 15, 15, and 60 seconds respectively
	// You can adjust these values as needed, or make them configurable
	// port also can be made configurable
	server := http.NewServer(15, 15, 60)

	// Register the HelloWorldHandler for the /hello-world endpoint
	server.RegisterHandler("/hello-world", handler.HelloWorldHandlerFunc())

	go server.Start("8080")

	<-quit // Wait for an interrupt signal
	log.Println("Shutdown signal received...")

	log.Println("Attempting graceful shutdown...")
	server.Stop()

	log.Println("Server gracefully stopped")
}
