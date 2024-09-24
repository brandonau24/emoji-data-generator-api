//nolint:errcheck
package request_handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	data_generation "github.com/brandonau24/emoji-data-generator/cmd/api_server/internal"
	"github.com/brandonau24/emoji-data-generator/cmd/api_server/internal/providers"
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

		urlProvider := providers.UnicodeDataUrlProvider{
			Version: version,
		}
		emojiDataGenerator := data_generation.EmojiDataGenerator{}
		emojis, err := emojiDataGenerator.Generate(urlProvider)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

			return
		}

		encodeError := encoder.Encode(emojis)

		if encodeError == nil {
			w.Header().Set("Content-Type", "application/json")
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
