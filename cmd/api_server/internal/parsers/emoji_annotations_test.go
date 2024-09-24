//nolint:errcheck
package parsers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseAnnotations(t *testing.T) {
	unicodeGithubTestServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() == "/unicode-org/cldr-json/refs/heads/main/cldr-json/cldr-annotations-modern/annotations/en/annotations.json" {
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

	defer unicodeGithubTestServer.Close()

	annotations := ParseAnnotations(unicodeGithubTestServer.URL)

	expectedAnnotations := []string{"face", "grin", "grinning face"}

	emojiAnnotations, ok := annotations["😀"]
	if !ok {
		t.Errorf("Failed to find annotations for 😀")
	}

	if !areAnnotationsEqual(emojiAnnotations.Default, expectedAnnotations) {
		t.Errorf("Failed to map annotations. Received %v, expected %v", emojiAnnotations, expectedAnnotations)
	}

	if emojiAnnotations.Tts[0] != "grinning face" {
		t.Errorf("Failed to map tts. Received %v, expected \"grinning face\"", emojiAnnotations.Tts)
	}
}
