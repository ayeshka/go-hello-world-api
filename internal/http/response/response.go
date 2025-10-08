package response

import (
	"encoding/json"
	"net/http"
)

// Response represents the JSON response structure
type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

// SuccessResponse creates a success response with a message
func SuccessResponse(message string) Response {
	return Response{Message: message}
}

// ErrorResponse creates an error response with an error message
func ErrorResponse(error string) Response {
	return Response{Error: error}
}

// Encoder handles JSON encoding of responses to HTTP response writers
type Encoder struct{}

// NewEncoder creates a new response encoder
func NewEncoder() *Encoder {
	return &Encoder{}
}

// EncodeSuccess writes a success response with the given message to the HTTP response writer
func (e *Encoder) EncodeSuccess(w http.ResponseWriter, message string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(SuccessResponse(message))
}

// EncodeError writes an error response with the given message and status code to the HTTP response writer
func (e *Encoder) EncodeError(w http.ResponseWriter, statusCode int, errorMessage string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(ErrorResponse(errorMessage))
}

// EncodeBadRequest writes a 400 Bad Request error response
func (e *Encoder) EncodeBadRequest(w http.ResponseWriter, errorMessage string) error {
	return e.EncodeError(w, http.StatusBadRequest, errorMessage)
}

// EncodeMethodNotAllowed writes a 405 Method Not Allowed error response
func (e *Encoder) EncodeMethodNotAllowed(w http.ResponseWriter, errorMessage string) error {
	return e.EncodeError(w, http.StatusMethodNotAllowed, errorMessage)
}

// Decoder handles JSON decoding from HTTP requests
type Decoder struct{}

// NewDecoder creates a new response decoder
func NewDecoder() *Decoder {
	return &Decoder{}
}

// DecodeResponse decodes a JSON response from an HTTP response body
func (d *Decoder) DecodeResponse(body []byte) (Response, error) {
	var resp Response
	err := json.Unmarshal(body, &resp)
	return resp, err
}

// WriteSuccessResponse is a convenience function to write a success response
func WriteSuccessResponse(w http.ResponseWriter, message string) error {
	encoder := NewEncoder()
	return encoder.EncodeSuccess(w, message)
}

// WriteBadRequestError is a convenience function to write a bad request error
func WriteBadRequestError(w http.ResponseWriter, errorMessage string) error {
	encoder := NewEncoder()
	return encoder.EncodeBadRequest(w, errorMessage)
}

// WriteMethodNotAllowedError is a convenience function to write a method not allowed error
func WriteMethodNotAllowedError(w http.ResponseWriter, errorMessage string) error {
	encoder := NewEncoder()
	return encoder.EncodeMethodNotAllowed(w, errorMessage)
}
