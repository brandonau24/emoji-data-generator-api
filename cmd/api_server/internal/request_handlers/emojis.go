//nolint:errcheck
package request_handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	data_generation "github.com/brandonau24/emoji-data-generator/cmd/api_server/internal"
	"github.com/brandonau24/emoji-data-generator/cmd/api_server/internal/providers"
)

type EmojisHandler struct{}

type EmojisRequestBody struct {
	Version string
}

func (h *EmojisHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var buffer bytes.Buffer
		encoder := json.NewEncoder(&buffer)
		encoder.SetEscapeHTML(false)

		defer r.Body.Close()
		requestBodyBytes, _ := io.ReadAll(r.Body)

		var requestBody EmojisRequestBody
		jsonErr := json.Unmarshal(requestBodyBytes, &requestBody)

		if jsonErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Request body contains Invalid JSON"))
		}

		_, parseError := strconv.ParseFloat(requestBody.Version, 64)

		if parseError != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("%v is not a number. The version should be a number e.g. 15.1"))

			return
		}

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
