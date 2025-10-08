# Simple HTTP API

A simple GoLang HTTP API with a single endpoint that validates names based on the first letter of the English alphabet.

## Overview

This service provides a single endpoint `/hello-world` that:
- Accepts a `name` query parameter
- Returns a greeting if the name starts with A-M (first half of alphabet)
- Returns an error if the name starts with N-Z (second half of alphabet)
- Returns an error for invalid or missing names

## Requirements

- Go 1.18 or later (tested with Go 1.21)

### Package Organization

- **`pkg/validation`** - Name validation logic (exportable package)
- **`internal/http/response`** - JSON encoding/decoding utilities (exportable package)
- **`internal/http/handler`** - HTTP handlers (internal package)
- **`main.go`** - Application entry point and server setup

## API Specification

### Endpoint: `/hello-world`

**Method:** GET

**Query Parameters:**
- `name` (required): The name to process

**Response Format:** JSON

#### Success Response (200 OK)
When the name starts with A-M (case-insensitive):
```json
{
  "message": "Hello {name}"
}
```

#### Error Response (400 Bad Request)
When the name starts with N-Z, is missing, empty, or invalid:
```json
{
  "error": "Invalid Input"
}
```

## Examples

### Valid Names (A-M)
```bash
curl "http://localhost:8080/hello-world?name=Alice"
# Response: 200 OK - {"message":"Hello Alice"}

curl "http://localhost:8080/hello-world?name=Alice%20Nam"
# Response: 200 OK - {"message":"Hello Alice Nam"}

curl "http://localhost:8080/hello-world?name=John%20Doe%20Jr"
# Response: 200 OK - {"message":"Hello John Doe Jr"}

curl "http://localhost:8080/hello-world?name=mark"
# Response: 200 OK - {"message":"Hello mark"}
```

### Invalid Names (N-Z)
```bash
curl "http://localhost:8080/hello-world?name=Nick"
# Response: 400 Bad Request - {"error":"Invalid Input"}

curl "http://localhost:8080/hello-world?name=Nancy%20Smith"
# Response: 400 Bad Request - {"error":"Invalid Input"}

curl "http://localhost:8080/hello-world?name=zoe"
# Response: 400 Bad Request - {"error":"Invalid Input"}
```

### Error Cases
```bash
curl "http://localhost:8080/hello-world"
# Response: 400 Bad Request - {"error":"Invalid Input"}

curl "http://localhost:8080/hello-world?name="
# Response: 400 Bad Request - {"error":"Invalid Input"}

curl "http://localhost:8080/hello-world?name=123John"
# Response: 400 Bad Request - {"error":"Invalid Input"}
```

## Running the Application

### Method 1: Direct Go Execution

1. **Clone or download the code to your workspace**

2. **Navigate to the project directory:**
   ```bash
   cd /go-hello-world-api
   ```

3. **Run the application:**
   ```bash
   go run .
   ```

   The server will start on port 8080 and log:
   ```
   Server starting on port 8080...
   Press Ctrl+C to gracefully shutdown the server
   ```

4. **Test the endpoint:**
   ```bash
   curl "http://localhost:8080/hello-world?name=Alice"
   ```

5. **Graceful shutdown:**
    - Press `Ctrl+C` or send `SIGTERM`/`SIGINT` signal
    - Server will complete active requests and shutdown gracefully
    - 30-second timeout for completion of active requests

## Running the Tests

Run all unit tests:
```bash
go test
```

Run tests with verbose output:
```bash
go test -v
```

Run tests with coverage:
```bash
go test -cover
```

Run tests with detailed coverage report:
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Design Decisions & Assumptions

1. **Alphabet Range**: The first half is defined as A-M (inclusive), second half as N-Z (inclusive)

2. **Case Sensitivity**: The validation is case-insensitive - both 'Alice' and 'alice' are treated the same

3. **Input Validation**:
    - Names must start with a letter (not numbers or special characters)
    - Empty strings and missing parameters are treated as invalid
    - Leading/trailing whitespace is preserved in the response but affects validation

4. **HTTP Methods**: Only GET requests are allowed, other methods return 405 Method Not Allowed

5. **Response Format**: All responses are in JSON format with appropriate HTTP status codes

6. **Error Handling**: All error cases return the same generic "Invalid Input" message as specified in requirements

7. **Multi-word Names**: Names with spaces are allowed (e.g., "John Doe") and the validation applies only to the first character

## Dependencies

This application uses only Go standard library packages:
- `net/http` - HTTP server and client
- `encoding/json` - JSON encoding/decoding
- `strings` - String manipulation
- `unicode` - Unicode character classification
- `log` - Logging