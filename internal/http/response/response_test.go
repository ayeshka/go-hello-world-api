package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccessResponse(t *testing.T) {
	message := "Hello World"
	resp := SuccessResponse(message)

	if resp.Message != message {
		t.Errorf("SuccessResponse() message = %q, want %q", resp.Message, message)
	}

	if resp.Error != "" {
		t.Errorf("SuccessResponse() error = %q, want empty string", resp.Error)
	}
}

func TestErrorResponse(t *testing.T) {
	errorMsg := "Invalid Input"
	resp := ErrorResponse(errorMsg)

	if resp.Error != errorMsg {
		t.Errorf("ErrorResponse() error = %q, want %q", resp.Error, errorMsg)
	}

	if resp.Message != "" {
		t.Errorf("ErrorResponse() message = %q, want empty string", resp.Message)
	}
}

func TestEncoder_EncodeSuccess(t *testing.T) {
	encoder := NewEncoder()
	rr := httptest.NewRecorder()
	message := "Hello Alice"

	err := encoder.EncodeSuccess(rr, message)
	if err != nil {
		t.Fatalf("EncodeSuccess() error = %v", err)
	}

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("EncodeSuccess() status = %v, want %v", status, http.StatusOK)
	}

	// Check content type
	expectedContentType := "application/json"
	if ct := rr.Header().Get("Content-Type"); ct != expectedContentType {
		t.Errorf("EncodeSuccess() Content-Type = %q, want %q", ct, expectedContentType)
	}

	// Check response body
	var resp Response
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("Could not unmarshal response: %v", err)
	}

	if resp.Message != message {
		t.Errorf("EncodeSuccess() response message = %q, want %q", resp.Message, message)
	}
}

func TestEncoder_EncodeError(t *testing.T) {
	encoder := NewEncoder()
	rr := httptest.NewRecorder()
	errorMsg := "Invalid Input"
	statusCode := http.StatusBadRequest

	err := encoder.EncodeError(rr, statusCode, errorMsg)
	if err != nil {
		t.Fatalf("EncodeError() error = %v", err)
	}

	// Check status code
	if status := rr.Code; status != statusCode {
		t.Errorf("EncodeError() status = %v, want %v", status, statusCode)
	}

	// Check content type
	expectedContentType := "application/json"
	if ct := rr.Header().Get("Content-Type"); ct != expectedContentType {
		t.Errorf("EncodeError() Content-Type = %q, want %q", ct, expectedContentType)
	}

	// Check response body
	var resp Response
	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("Could not unmarshal response: %v", err)
	}

	if resp.Error != errorMsg {
		t.Errorf("EncodeError() response error = %q, want %q", resp.Error, errorMsg)
	}
}

func TestEncoder_EncodeBadRequest(t *testing.T) {
	encoder := NewEncoder()
	rr := httptest.NewRecorder()
	errorMsg := "Bad Request"

	err := encoder.EncodeBadRequest(rr, errorMsg)
	if err != nil {
		t.Fatalf("EncodeBadRequest() error = %v", err)
	}

	// Check status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("EncodeBadRequest() status = %v, want %v", status, http.StatusBadRequest)
	}
}

func TestEncoder_EncodeMethodNotAllowed(t *testing.T) {
	encoder := NewEncoder()
	rr := httptest.NewRecorder()
	errorMsg := "Method not allowed"

	err := encoder.EncodeMethodNotAllowed(rr, errorMsg)
	if err != nil {
		t.Fatalf("EncodeMethodNotAllowed() error = %v", err)
	}

	// Check status code
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("EncodeMethodNotAllowed() status = %v, want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestDecoder_DecodeResponse(t *testing.T) {
	decoder := NewDecoder()

	tests := []struct {
		name     string
		jsonData string
		expected Response
		wantErr  bool
	}{
		{
			name:     "Valid success response",
			jsonData: `{"message":"Hello World"}`,
			expected: Response{Message: "Hello World"},
			wantErr:  false,
		},
		{
			name:     "Valid error response",
			jsonData: `{"error":"Invalid Input"}`,
			expected: Response{Error: "Invalid Input"},
			wantErr:  false,
		},
		{
			name:     "Invalid JSON",
			jsonData: `{invalid json}`,
			expected: Response{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := decoder.DecodeResponse([]byte(tt.jsonData))

			if tt.wantErr && err == nil {
				t.Errorf("DecodeResponse() expected error but got none")
			}

			if !tt.wantErr && err != nil {
				t.Errorf("DecodeResponse() unexpected error = %v", err)
			}

			if !tt.wantErr && (resp.Message != tt.expected.Message || resp.Error != tt.expected.Error) {
				t.Errorf("DecodeResponse() = %+v, want %+v", resp, tt.expected)
			}
		})
	}
}

func TestWriteSuccessResponse(t *testing.T) {
	rr := httptest.NewRecorder()
	message := "Hello Test"

	err := WriteSuccessResponse(rr, message)
	if err != nil {
		t.Fatalf("WriteSuccessResponse() error = %v", err)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("WriteSuccessResponse() status = %v, want %v", status, http.StatusOK)
	}
}

func TestWriteBadRequestError(t *testing.T) {
	rr := httptest.NewRecorder()
	errorMsg := "Bad Request Test"

	err := WriteBadRequestError(rr, errorMsg)
	if err != nil {
		t.Fatalf("WriteBadRequestError() error = %v", err)
	}

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("WriteBadRequestError() status = %v, want %v", status, http.StatusBadRequest)
	}
}

func TestWriteMethodNotAllowedError(t *testing.T) {
	rr := httptest.NewRecorder()
	errorMsg := "Method Not Allowed Test"

	err := WriteMethodNotAllowedError(rr, errorMsg)
	if err != nil {
		t.Fatalf("WriteMethodNotAllowedError() error = %v", err)
	}

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("WriteMethodNotAllowedError() status = %v, want %v", status, http.StatusMethodNotAllowed)
	}
}
