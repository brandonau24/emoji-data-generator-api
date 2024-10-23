package request_handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestHandler(t *testing.T) {
	tests := map[string]string{
		"rejects Connect": http.MethodConnect,
		"rejects Delete":  http.MethodDelete,
		"rejects Head":    http.MethodHead,
		"rejects Patch":   http.MethodPatch,
		"rejects Post":    http.MethodPost,
		"rejects Put":     http.MethodPut,
		"rejects Trace":   http.MethodTrace,
	}

	for name, httpMethod := range tests {
		t.Run(name, func(t *testing.T) {
			emojiHandler := EmojisHandler{}

			request := httptest.NewRequest(httpMethod, "/", nil)
			responseRecorder := httptest.NewRecorder()

			emojiHandler.ServeHTTP(responseRecorder, request)

			response := responseRecorder.Result()
			defer response.Body.Close()

			if response.StatusCode != http.StatusMethodNotAllowed {
				t.Errorf("Expected %v to respond with %v, got: %v", httpMethod, http.StatusMethodNotAllowed, response.StatusCode)
			}
		})
	}
}

func TestRequestHandler_WithQueryParameters(t *testing.T) {
	tests := map[string]string{
		"rejects non-number version":                "abcdef",
		"rejects mixed numbers and letters version": "12ab3cd456",
	}

	for name, queryParam := range tests {
		t.Run(name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/?unicode_version=%v", queryParam), nil)
			responseRecorder := httptest.NewRecorder()

			emojiHandler := EmojisHandler{}
			emojiHandler.ServeHTTP(responseRecorder, request)

			response := responseRecorder.Result()
			defer response.Body.Close()

			if response.StatusCode != http.StatusBadRequest {
				t.Errorf("Expected 400 status code, received %v", response.StatusCode)
			}
		})
	}
}

// TODO: This is more of an integration/end-to-end test since it is testing the real connection without mocks. Need to move it to a health check
// func TestApiAllowsEmptyRequestBody(t *testing.T) {
// 	request := httptest.NewRequest(http.MethodGet, "/", nil)
// 	responseRecorder := httptest.NewRecorder()

// 	emojiHandler := EmojisHandler{}
// 	emojiHandler.ServeHTTP(responseRecorder, request)

// 	response := responseRecorder.Result()
// 	defer response.Body.Close()

// 	if response.StatusCode != http.StatusOK {
// 		t.Errorf("Expected 200 status code, received %v", response.StatusCode)
// 	}
// }
