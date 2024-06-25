package parsers_test

import (
	"testing"

	"github.com/brandonau24/emoji-data-generator/parsers"
)

func areAnnotationsEqual(annotations1 []string, annotations2 []string) bool {
	if len(annotations1) != len(annotations2) {
		return false
	}

	for i := range annotations1 {
		if annotations1[i] != annotations2[i] {
			return false
		}
	}

	return true
}

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

	if !areAnnotationsEqual(emojiAnnotations, expectedAnnotations) {
		t.Fatalf("Failed to map annotations. Received %v, expected %v", emojiAnnotations, expectedAnnotations)
	}
}
