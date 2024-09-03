package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brandonau24/emoji-data-generator/parsers"
)

func TestEmojiParserApi(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)

	emojiHandler := EmojiHandler{}
	responseRecorder := httptest.NewRecorder()

	emojiHandler.ServeHTTP(responseRecorder, request)

	response := responseRecorder.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected 200, received %v", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)

	if err != nil {
		t.Errorf("Unexpected error in reading response body: %v", err)
	}

	var emojiData map[string][]parsers.Emoji
	json.Unmarshal(data, &emojiData)

	// TODO: Need stronger assertions. Check for each group? Check a few emojis and their properties?
	if _, ok := emojiData["Smileys & Emotion"]; !ok {
		t.Errorf("Expected \"Smileys & Emotion\" group to exist, but it does not: %v", emojiData)
	}
}
