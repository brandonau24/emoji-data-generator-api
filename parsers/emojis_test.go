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
	emojis := parsers.ParseEmojis(`# group: group1
1F600                                                  ; fully-qualified     # 😀 E1.0 grinning face`)
	emoji := emojis["group1"][0]

	if emoji.Codepoints != "1F600" {
		t.Fatalf("Failed to parse codepoint. Received %v, expected 1F600", emoji.Codepoints)
	}
}

func TestParseEmojisSetsCodepoints(t *testing.T) {
	emojis := parsers.ParseEmojis(` # group: group1
1F62E 200D 1F4A8                                       ; fully-qualified     # 😮‍💨 E13.1 face exhaling
`)
	emoji := emojis["group1"][0]

	if emoji.Codepoints != "1F62E 200D 1F4A8" {
		t.Fatalf("Failed to parse codepoint. Received %v, expected 1F62E 200D 1F4A8", emoji.Codepoints)
	}
}

func TestParseEmojisSetsName(t *testing.T) {
	emojis := parsers.ParseEmojis(`# group: group1
1F600                                                  ; fully-qualified     # 😀 E1.0 grinning face`)
	emoji := emojis["group1"][0]

	if emoji.Name != "grinning face" {
		t.Fatalf("Failed to parse codepoint. Received %v, expected grinning face", emoji.Name)
	}
}

func TestParseEmojiSelectsFullyQualifiedEmojis(t *testing.T) {
	emojis := parsers.ParseEmojis(` # group: group1
F636 200D 1F32B FE0F                                  ; fully-qualified     # 😶‍🌫️ E13.1 face in clouds
1F636 200D 1F32B                                       ; minimally-qualified # 😶‍🌫 E13.1 face in clouds
2620                                                   ; unqualified         # ☠ E1.0 skull and crossbones
`)
	emojisInGroup1 := emojis["group1"]

	if len(emojisInGroup1) != 1 {
		t.Fatalf("Other emojis that are not fully qualified have been parsed")
	}
}

func TestParseEmojisGroupsEmojis(t *testing.T) {
	emojis := parsers.ParseEmojis(`# group: Smileys & Emotion
1F600                                                  ; fully-qualified     # 😀 E1.0 grinning face
1F603                                                  ; fully-qualified     # 😃 E0.6 grinning face with big eyes
1F636 200D 1F32B                                       ; minimally-qualified # 😶‍🌫 E13.1 face in clouds

# group: People & Body
1F44B                                                  ; fully-qualified     # 👋 E0.6 waving hand
1F44B 1F3FB                                            ; fully-qualified     # 👋🏻 E1.0 waving hand: light skin tone
1F590                                                  ; unqualified         # 🖐 E0.7 hand with fingers splayed
`)

	smileyAndEmotionsGroup, smileyAndEmotionOk := emojis["Smileys & Emotion"]

	if !smileyAndEmotionOk {
		t.Fatalf("Could not parse Smileys & Emotion group")
	}

	if len(smileyAndEmotionsGroup) != 2 {
		t.Fatalf("Expected 2 emojis in Smiley & Emotion group, received %v", len(smileyAndEmotionsGroup))
	}

	peopleAndBodyGroup, peopleAndBodyOk := emojis["People & Body"]

	if !peopleAndBodyOk {
		t.Fatalf("Could not parse People & Body group")
	}

	if len(peopleAndBodyGroup) != 2 {
		t.Fatalf("Expected 2 emojis in People & Body group, received %v", len(peopleAndBodyGroup))
	}
}
