//nolint:errcheck
package request_handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	data_generation "github.com/brandonau24/emoji-data-generator/cmd/api_server/internal"
	"github.com/brandonau24/emoji-data-generator/cmd/api_server/internal/providers"
)

type EmojisHandler struct{}

func (h *EmojisHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var buffer bytes.Buffer
		encoder := json.NewEncoder(&buffer)
		encoder.SetEscapeHTML(false)

		urlProvider := providers.UnicodeDataUrlProvider{}
		emojiDataGenerator := data_generation.EmojiDataGenerator{}
		emojis, err := emojiDataGenerator.Generate(urlProvider)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

			return
		}

		encodeError := encoder.Encode(emojis)

		if encodeError == nil {
			w.Write(buffer.Bytes())
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Could not parse emoji data"))
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprintf("%v request not allowed", r.Method)))
	}
}
