package parsers_test

import (
	"testing"

	"github.com/brandonau24/emoji-data-generator/parsers"
)

func TestParseEmojisSkipsComments(t *testing.T) {
	emojis := parsers.ParseEmojis(`
		# This is a comment
		# This is another comment
		# This is the last comment
	`)

	if len(emojis) != 0 {
		t.Fatalf("Failed to parse comments")
	}
}
