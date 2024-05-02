package parsers_test

import (
	"testing"

	"github.com/brandonau24/emoji-data-generator/parsers"
)

func TestParseEmojisSkipsComments(t *testing.T) {
	emojis := parsers.ParseEmojis(`# This is a comment
# This is another comment
# This is the last comment`)

	if len(emojis) != 0 {
		t.Fatalf("Failed to parse comments")
	}
}

func TestParseEmojisSetsCodepoint(t *testing.T) {
	emojis := parsers.ParseEmojis("1F600                                                  ; fully-qualified     # 😀 E1.0 grinning face")
	emoji := emojis[0]

	if emoji.Codepoints != "1F600" {
		t.Fatalf("Failed to parse codepoint. Received %v, expected 1F600", emoji.Codepoints)
	}
}
