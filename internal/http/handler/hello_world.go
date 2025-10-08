package handler

import (
	"go-hello-world-api/internal/http/response"
	"go-hello-world-api/pkg/validation"
	"net/http"
)

// HelloWorldHandler handles the /hello-world endpoint
type HelloWorldHandler struct {
	encoder *response.Encoder
}

// NewHelloWorldHandler creates a new HelloWorldHandler with the given encoder
func NewHelloWorldHandler() *HelloWorldHandler {
	return &HelloWorldHandler{
		encoder: response.NewEncoder(),
	}
}

// ServeHTTP implements the http.Handler interface for the HelloWorldHandler
func (h *HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		h.encoder.EncodeMethodNotAllowed(w, "Method not allowed")
		return
	}

	// Get the name parameter from query string
	name := r.URL.Query().Get("name")

	// Validate the name using our validation function
	if validation.ValidateName(name) {

		h.encoder.EncodeSuccess(w, "Hello "+name)
	} else {

		h.encoder.EncodeBadRequest(w, "Invalid Input")
	}
}

func HelloWorldHandlerFunc() http.HandlerFunc {
	handler := NewHelloWorldHandler()
	return handler.ServeHTTP
}
