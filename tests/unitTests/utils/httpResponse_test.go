package utils

import (
	"httpServer/src/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpResponse(t *testing.T) {
	// Define test cases
	tests := []struct {
		name           string
		inputData      interface{}
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Success response with JSON",
			inputData:      map[string]string{"message": "success"},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"success"}`,
		},
		{
			name:           "Internal server error on JSON marshal failure",
			inputData:      make(chan int),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Failed to marshal response to JSON\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			utils.HttpResponse(tt.inputData, rr, tt.expectedStatus)
			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
			if rr.Body.String() != tt.expectedBody {
				t.Errorf("Expected body %s, got %s", tt.expectedBody, rr.Body.String())
			}
			if rr.Header().Get("Content-Type") != "application/json" && tt.expectedStatus != http.StatusInternalServerError {
				t.Errorf("Expected Content-Type application/json, got %s", rr.Header().Get("Content-Type"))
			}
		})
	}
}
