package parsers

import (
	"testing"
)

func TestParseAnnotations(t *testing.T) {
	annotations := ParseAnnotations(`
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
