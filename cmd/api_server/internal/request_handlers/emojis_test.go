package request_handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiOnlyAcceptsGetRequest(t *testing.T) {
	unallowedHttpMethods := []string{http.MethodConnect, http.MethodDelete, http.MethodPost, http.MethodHead, http.MethodPatch, http.MethodDelete, http.MethodPut, http.MethodTrace}
	emojiHandler := EmojisHandler{}

	for _, method := range unallowedHttpMethods {
		request := httptest.NewRequest(method, "/", nil)
		responseRecorder := httptest.NewRecorder()

		emojiHandler.ServeHTTP(responseRecorder, request)

		response := responseRecorder.Result()
		defer response.Body.Close()

		if response.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Expected 405 status code, received %v", response.StatusCode)
		}
	}
}
