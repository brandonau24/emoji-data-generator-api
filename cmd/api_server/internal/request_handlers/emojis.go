//nolint:errcheck
package request_handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

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

		defer r.Body.Close()
		requestBodyBytes, _ := io.ReadAll(r.Body)

		var version float64
		if len(requestBodyBytes) > 0 {
			var requestBody EmojisRequestBody
			jsonErr := json.Unmarshal(requestBodyBytes, &requestBody)

			if jsonErr != nil {
				log.Println(jsonErr.Error())
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Request body contains Invalid JSON"))

				return
			}

			version = requestBody.Version
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
