package parsers

import (
	"testing"
)

func Test_ParseCodepoint(t *testing.T) {
	emojiFields := []string{"1F600", ";", "fully-qualified", "#", "😀", "E1.0", "grinning face"}

	codepoints := ParseCodepoints(emojiFields)

	if codepoints != "1F600" {
		t.Errorf("Expected 1F600 codepoint, received %v", codepoints)
	}
}

func Test_ParseCodepoints(t *testing.T) {
	emojiFields := []string{"1F636", "200D", "1F32B", "FE0F", ";", "fully-qualified", "#", "😶‍🌫️", "E13.1", "face", "in", "clouds"}

	codepoints := ParseCodepoints(emojiFields)

	if codepoints != "1F636 200D 1F32B FE0F" {
		t.Errorf("Expected 1F636 200D 1F32B FE0F codepoints, received %v", codepoints)
	}
}

func Test_ParseEmojiName(t *testing.T) {
	emojiFields := []string{"1F600", ";", "fully-qualified", "#", "😀", "E1.0", "grinning face"}

	name := ParseEmojiName(emojiFields)

	if name != "grinning face" {
		t.Errorf("Expected \"grinning face\", received %v", name)
	}
}

func Test_ParseEmojiCharacter(t *testing.T) {
	emojiFields := []string{"1F600", ";", "fully-qualified", "#", "😀", "E1.0", "grinning face"}

	character := ParseEmojiCharacter(emojiFields)

	if character != "😀" {
		t.Errorf("Expected \"😀\", received %v", character)
	}
}
