package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/brandonau24/emoji-data-generator/cmd/api_server/internal/parsers"
)

type EmojiHandler struct{}

func (h *EmojiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var buffer bytes.Buffer
		encoder := json.NewEncoder(&buffer)
		encoder.SetEscapeHTML(false)

		emojiAnnotations := parsers.ParseAnnotations("")

		emojis, parseError := parsers.ParseEmojis("", emojiAnnotations)

		if parseError != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(parseError.Error())) //nolint:errcheck

			return
		}

		encodeError := encoder.Encode(emojis)

		if encodeError == nil {
			w.Write(buffer.Bytes()) //nolint:errcheck
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Could not parse emoji data")) //nolint:errcheck
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprintf("%v request not allowed", r.Method))) //nolint:errcheck
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &EmojiHandler{})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
