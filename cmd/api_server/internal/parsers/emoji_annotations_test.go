//nolint:errcheck
package parsers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	test_helpers "github.com/brandonau24/emoji-data-generator-api/cmd/api_server/internal/internal"
)

func Test_ParseAnnotations(t *testing.T) {
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == test_helpers.MOCK_UNICODE_ANNOTATIONS_PATH {
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

	annotations := ParseAnnotations(test_helpers.MockDataUrlProvider{
		BaseUrl: mockHttpServer.URL,
	})

	expectedAnnotations := []string{"face", "grin", "grinning face"}

	emojiAnnotations, ok := annotations["😀"]
	if !ok {
		t.Errorf("Failed to find annotations for 😀")
	}

	if !test_helpers.AreAnnotationsEqual(emojiAnnotations.Default, expectedAnnotations) {
		t.Errorf("Failed to map annotations. Received %v, expected %v", emojiAnnotations, expectedAnnotations)
	}

	if emojiAnnotations.Tts[0] != "grinning face" {
		t.Errorf("Failed to map tts. Received %v, expected \"grinning face\"", emojiAnnotations.Tts)
	}
}

func Test_ParseAnnotations_ReturnsNil_WhenRequestFails(t *testing.T) {
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == test_helpers.MOCK_UNICODE_ANNOTATIONS_PATH {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(""))
		}
	}))

	defer mockHttpServer.Close()

	annotations := ParseAnnotations(test_helpers.MockDataUrlProvider{
		BaseUrl: mockHttpServer.URL,
	})

	if annotations != nil {
		t.Errorf("Expected no annotations returned when annotations request fails. Received %v", annotations)
	}
}
