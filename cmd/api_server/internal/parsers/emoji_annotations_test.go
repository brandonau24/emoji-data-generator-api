//nolint:errcheck
package parsers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	parsers_tests "github.com/brandonau24/emoji-data-generator/cmd/api_server/internal/parsers/internal"
)

func TestParseAnnotations(t *testing.T) {
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == parsers_tests.MOCK_UNICODE_ANNOTATIONS_PATH {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
{
	"annotations": {
		"annotations": {
			"😀": {
				"default": [
					"face",
					"grin",
					"grinning face"
				],
				"tts": [
					"grinning face"
				]
			}
		}
	}
}
`))
		}
	}))

	defer mockHttpServer.Close()

	annotations := ParseAnnotations(parsers_tests.MockDataUrlProvider{
		BaseUrl: mockHttpServer.URL,
	})

	expectedAnnotations := []string{"face", "grin", "grinning face"}

	emojiAnnotations, ok := annotations["😀"]
	if !ok {
		t.Errorf("Failed to find annotations for 😀")
	}

	if !parsers_tests.AreAnnotationsEqual(emojiAnnotations.Default, expectedAnnotations) {
		t.Errorf("Failed to map annotations. Received %v, expected %v", emojiAnnotations, expectedAnnotations)
	}

	if emojiAnnotations.Tts[0] != "grinning face" {
		t.Errorf("Failed to map tts. Received %v, expected \"grinning face\"", emojiAnnotations.Tts)
	}
}
