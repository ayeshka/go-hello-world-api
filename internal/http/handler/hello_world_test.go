package handler

import (
	"encoding/json"
	"go-hello-world-api/internal/http/response"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorldHandler_ServeHTTP(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		queryParam     string
		expectedStatus int
		expectedBody   response.Response
	}{
		{
			name:           "Valid name starting with A",
			method:         "GET",
			queryParam:     "Alice",
			expectedStatus: http.StatusOK,
			expectedBody:   response.Response{Message: "Hello Alice"},
		},
		{
			name:           "Valid name starting with a (lowercase)",
			method:         "GET",
			queryParam:     "alice",
			expectedStatus: http.StatusOK,
			expectedBody:   response.Response{Message: "Hello alice"},
		},
		{
			name:           "Valid name starting with M",
			method:         "GET",
			queryParam:     "Mark",
			expectedStatus: http.StatusOK,
			expectedBody:   response.Response{Message: "Hello Mark"},
		},
		{
			name:           "Valid name starting with m (lowercase)",
			method:         "GET",
			queryParam:     "mary",
			expectedStatus: http.StatusOK,
			expectedBody:   response.Response{Message: "Hello mary"},
		},
		{
			name:           "Valid name with multiple words",
			method:         "GET",
			queryParam:     "John Doe",
			expectedStatus: http.StatusOK,
			expectedBody:   response.Response{Message: "Hello John Doe"},
		},
		{
			name:           "Valid multi-word name Alice Nam",
			method:         "GET",
			queryParam:     "Alice Nam",
			expectedStatus: http.StatusOK,
			expectedBody:   response.Response{Message: "Hello Alice Nam"},
		},
		{
			name:           "Valid lowercase multi-word alice nam",
			method:         "GET",
			queryParam:     "alice nam",
			expectedStatus: http.StatusOK,
			expectedBody:   response.Response{Message: "Hello alice nam"},
		},
		{
			name:           "Valid multi-word with middle initial",
			method:         "GET",
			queryParam:     "Mary J Watson",
			expectedStatus: http.StatusOK,
			expectedBody:   response.Response{Message: "Hello Mary J Watson"},
		},
		{
			name:           "Valid hyphenated multi-word name",
			method:         "GET",
			queryParam:     "Anne-Marie Johnson",
			expectedStatus: http.StatusOK,
			expectedBody:   response.Response{Message: "Hello Anne-Marie Johnson"},
		},

		// Invalid names (N-Z)
		{
			name:           "Invalid name starting with N",
			method:         "GET",
			queryParam:     "Nick",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   response.Response{Error: "Invalid Input"},
		},
		{
			name:           "Invalid name starting with n (lowercase)",
			method:         "GET",
			queryParam:     "nick",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   response.Response{Error: "Invalid Input"},
		},
		{
			name:           "Invalid name starting with Z",
			method:         "GET",
			queryParam:     "Zane",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   response.Response{Error: "Invalid Input"},
		},
		{
			name:           "Invalid name starting with z (lowercase)",
			method:         "GET",
			queryParam:     "zoe",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   response.Response{Error: "Invalid Input"},
		},
		{
			name:           "Invalid multi-word starting with N",
			method:         "GET",
			queryParam:     "Nancy Smith",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   response.Response{Error: "Invalid Input"},
		},
		{
			name:           "Invalid multi-word starting with Z",
			method:         "GET",
			queryParam:     "Zoe Johnson",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   response.Response{Error: "Invalid Input"},
		},
		{
			name:           "Invalid lowercase multi-word starting with n",
			method:         "GET",
			queryParam:     "nancy smith",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   response.Response{Error: "Invalid Input"},
		},
		{
			name:           "Invalid multi-word starting with P",
			method:         "GET",
			queryParam:     "Peter Parker",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   response.Response{Error: "Invalid Input"},
		},

		// Edge cases
		{
			name:           "Empty name parameter",
			method:         "GET",
			queryParam:     "",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   response.Response{Error: "Invalid Input"},
		},
		{
			name:           "Name starting with number",
			method:         "GET",
			queryParam:     "123John",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   response.Response{Error: "Invalid Input"},
		},
		{
			name:           "Name starting with special character",
			method:         "GET",
			queryParam:     "@Alice",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   response.Response{Error: "Invalid Input"},
		},
		{
			name:           "Single character valid name",
			method:         "GET",
			queryParam:     "A",
			expectedStatus: http.StatusOK,
			expectedBody:   response.Response{Message: "Hello A"},
		},
		{
			name:           "Single character invalid name",
			method:         "GET",
			queryParam:     "Z",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   response.Response{Error: "Invalid Input"},
		},
		{
			name:           "Name with spaces at beginning",
			method:         "GET",
			queryParam:     " Alice",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   response.Response{Error: "Invalid Input"},
		},

		// Method tests
		{
			name:           "POST method not allowed",
			method:         "POST",
			queryParam:     "Alice",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   response.Response{Error: "Method not allowed"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request
			req, err := http.NewRequest(tt.method, "/hello-world?name="+tt.queryParam, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create response recorder
			rr := httptest.NewRecorder()

			// Create handler and call it
			handler := NewHelloWorldHandler()
			handler.ServeHTTP(rr, req)

			// Check status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			// Check response body
			var response response.Response
			err = json.Unmarshal(rr.Body.Bytes(), &response)
			if err != nil {
				t.Fatalf("Could not unmarshal response: %v", err)
			}

			if response.Message != tt.expectedBody.Message || response.Error != tt.expectedBody.Error {
				t.Errorf("handler returned unexpected body: got %v want %v", response, tt.expectedBody)
			}
		})
	}
}

func TestHelloWorldHandler_MissingParameter(t *testing.T) {
	// Test when name parameter is completely missing
	req, err := http.NewRequest("GET", "/hello-world", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := NewHelloWorldHandler()
	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	// Check response body
	var response response.Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Could not unmarshal response: %v", err)
	}

	expectedResponse := "Invalid Input"
	if response.Error != expectedResponse {
		t.Errorf("handler returned unexpected body: got %v want %v", response, expectedResponse)
	}
}

func TestHelloWorldHandlerFunc(t *testing.T) {
	// Test the convenience function
	handlerFunc := HelloWorldHandlerFunc()

	req, err := http.NewRequest("GET", "/hello-world?name=Alice", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handlerFunc(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response body
	var response response.Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Could not unmarshal response: %v", err)
	}

	expectedResponse := "Hello Alice"
	if response.Message != expectedResponse {
		t.Errorf("handler returned unexpected body: got %v want %v", response, expectedResponse)
	}
}
