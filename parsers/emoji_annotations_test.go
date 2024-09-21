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
		t.Errorf("Failed to find annotations for 1F600")
	}

	if !areAnnotationsEqual(emojiAnnotations.Default, expectedAnnotations) {
		t.Errorf("Failed to map annotations. Received %v, expected %v", emojiAnnotations, expectedAnnotations)
	}

	if emojiAnnotations.Tts[0] != "grinning face" {
		t.Errorf("Failed to map tts. Received %v, expected \"grinning face\"", emojiAnnotations.Tts)
	}
}
