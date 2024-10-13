//nolint:errcheck
package request_handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	data_generation "github.com/brandonau24/emoji-data-generator-api/cmd/api_server/internal"
	"github.com/brandonau24/emoji-data-generator-api/cmd/api_server/internal/providers"
)

type EmojisHandler struct{}

type EmojisRequestBody struct {
	Version float64
}

func (h *EmojisHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var buffer bytes.Buffer
		encoder := json.NewEncoder(&buffer)
		encoder.SetEscapeHTML(false)

		queryParams := r.URL.Query()
		unicodeVersion := queryParams.Get("unicode_version")
		version, parseErr := strconv.ParseFloat(unicodeVersion, 32)

		if unicodeVersion != "" && parseErr != nil {
			log.Println(parseErr.Error())
			http.Error(w, fmt.Sprintf("\"%v\" is not a valid Unicode version", unicodeVersion), http.StatusBadRequest)

			return
		}

		urlProvider := providers.UnicodeDataUrlProvider{}
		emojiDataGenerator := data_generation.EmojiDataGenerator{
			UrlProvider: urlProvider,
		}
		emojis, err := emojiDataGenerator.Generate(version)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		encodeError := encoder.Encode(emojis)

		if encodeError == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write(buffer.Bytes())

			return
		} else {
			http.Error(w, "Could not parse emoji data", http.StatusInternalServerError)

			return
		}
	} else {
		methodNotAllowedErrorMessage := fmt.Sprintf("%v request not allowed", r.Method)

		http.Error(w, methodNotAllowedErrorMessage, http.StatusMethodNotAllowed)

		return
	}
}
