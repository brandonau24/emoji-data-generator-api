package parsers

import (
	"testing"
)

func TestParseCodepoints(t *testing.T) {
	tests := map[string]struct {
		expectedCodepoints string
		emojiFields        []string
	}{
		"parses a single codepoint": {
			expectedCodepoints: "1F600",
			emojiFields:        []string{"1F600", ";", "fully-qualified", "#", "😀", "E1.0", "grinning", "face"},
		},
		"parses multiple codepoints": {
			expectedCodepoints: "1F636 200D 1F32B FE0F",
			emojiFields:        []string{"1F636", "200D", "1F32B", "FE0F", ";", "fully-qualified", "#", "😶‍🌫️", "E13.1", "face", "in", "clouds"},
		},
	}

	for name, testcase := range tests {
		t.Run(name, func(t *testing.T) {
			codepoints := ParseCodepoints(testcase.emojiFields)

			if codepoints != testcase.expectedCodepoints {
				t.Errorf("%v: expected: %v, got: %v", name, testcase.expectedCodepoints, codepoints)
			}
		})
	}
}

func TestParseEmojiName(t *testing.T) {
	tests := map[string]struct {
		expectedName string
		emojiFields  []string
	}{
		"parses emoji name containing one word": {
			expectedName: "ogre",
			emojiFields:  []string{"1F479", ";", "fully-qualified", "#", "👹", "E0.6", "ogre"},
		},
		"parses emoji name containing multiple words": {
			expectedName: "skull and crossbones",
			emojiFields:  []string{"2620", "FE0F", ";", "fully-qualified", "#", "☠️", "E1.0", "skull", "and", "crossbones"},
		},
	}

	for name, testcase := range tests {
		t.Run(name, func(t *testing.T) {
			emojiName := ParseEmojiName(testcase.emojiFields)

			if emojiName != testcase.expectedName {
				t.Errorf("%v: expected: %v, got: %v", name, testcase.expectedName, emojiName)
			}
		})
	}
}

func TestParseEmojiCharacter(t *testing.T) {
	tests := map[string]struct {
		expectedEmoji string
		emojiFields   []string
	}{
		"parses emoji character": {
			expectedEmoji: "😀",
			emojiFields:   []string{"1F600", ";", "fully-qualified", "#", "😀", "E1.0", "grinning face"},
		},
		"parses multicomponent emoji character": {
			expectedEmoji: "⛓️‍💥",
			emojiFields:   []string{"26D3", "FE0F", "200D", "1F4A5", "fully-qualified", "#", "⛓️‍💥", "E15.1", "broken", "chain"}, // Broken chain is actually ⛓️‍ + 💥
		},
	}

	for name, testcase := range tests {
		t.Run(name, func(t *testing.T) {
			emoji := ParseEmojiCharacter(testcase.emojiFields)

			if emoji != testcase.expectedEmoji {
				t.Errorf("%v: expected: %v, got: %v", name, testcase.expectedEmoji, emoji)
			}
		})
	}
}
