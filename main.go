package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/brandonau24/emoji-data-generator/parsers"
	"github.com/brandonau24/emoji-data-generator/readers"
)

type EmojiHandler struct{}

func (h *EmojiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var buffer bytes.Buffer
		encoder := json.NewEncoder(&buffer)
		encoder.SetEscapeHTML(false)

		emojiDataFile := readers.ReadEmojiDataFile()
		emojiAnnotationsFile := readers.ReadEmojiAnnotationsFile()

		emojiAnnotations := parsers.ParseAnnotations(emojiAnnotationsFile)

		emojis := parsers.ParseEmojis(emojiDataFile, emojiAnnotations)
		err := encoder.Encode(emojis)

		if err == nil {
			w.Write(buffer.Bytes())
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Could not parse emoji data"))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad request"))
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &EmojiHandler{})

	http.ListenAndServe(":8080", mux)
}
