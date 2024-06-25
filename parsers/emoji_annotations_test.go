package parsers_test

import (
	"testing"

	"github.com/brandonau24/emoji-data-generator/parsers"
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

	emojiAnnotations, ok := annotations["1F600"]
	if !ok {
		t.Fatalf("Failed to find annotations for 1F600")
	}

	if len(emojiAnnotations) != 3 {
		t.Fatalf("Failed to set annotations for 1F600")
	}
}
