package parsers_test

import (
	"testing"

	"github.com/brandonau24/emoji-data-generator/parsers"
	"github.com/brandonau24/emoji-data-generator/test_helpers"
)

func TestParseAnnotations(t *testing.T) {
	annotations := parsers.ParseAnnotations(`
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
`)

	expectedAnnotations := []string{"face", "grin", "grinning face"}

	emojiAnnotations, ok := annotations["1F600"]
	if !ok {
		t.Fatalf("Failed to find annotations for 1F600")
	}

	if !test_helpers.AreAnnotationsEqual(emojiAnnotations, expectedAnnotations) {
		t.Fatalf("Failed to map annotations. Received %v, expected %v", emojiAnnotations, expectedAnnotations)
	}
}
