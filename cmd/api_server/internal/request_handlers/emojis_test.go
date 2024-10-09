package request_handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Api_OnlyAcceptsGetRequest(t *testing.T) {
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

func Test_Api_RejectsNonNumberVersion(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/?unicode_version=abcdef", nil)
	responseRecorder := httptest.NewRecorder()

	emojiHandler := EmojisHandler{}
	emojiHandler.ServeHTTP(responseRecorder, request)

	response := responseRecorder.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected 400 status code, received %v", response.StatusCode)
	}
}
