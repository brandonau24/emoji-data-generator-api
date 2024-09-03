package main

import (
	"encoding/json"
	"net/http"

	"github.com/brandonau24/emoji-data-generator/parsers"
	"github.com/brandonau24/emoji-data-generator/readers"
)

type EmojiHandler struct{}

func (h *EmojiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		emojiDataFile := readers.ReadEmojiDataFile()
		emojiAnnotationsFile := readers.ReadEmojiAnnotationsFile()

		emojiAnnotations := parsers.ParseAnnotations(emojiAnnotationsFile)

		emojis := parsers.ParseEmojis(emojiDataFile, emojiAnnotations)
		emojisJson, err := json.Marshal(emojis)

		if err == nil {
			w.Write(emojisJson)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Could not parse emoji data"))
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &EmojiHandler{})

	http.ListenAndServe(":8080", mux)
}
