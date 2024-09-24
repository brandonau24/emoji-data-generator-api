package parsers

import (
	"regexp"
	"slices"
	"strings"
)

func ParseCodepoints(emojiFields []string) string {
	semicolonIndex := slices.Index(emojiFields, ";")
	codepoints := emojiFields[0:semicolonIndex]

	return strings.Join(codepoints, " ")
}

func ParseEmojiName(emojiFields []string) string {
	emojiVersionRegex := regexp.MustCompile(`E\d+\.\d+`)
	emojiVersionIndex := slices.IndexFunc(emojiFields, func(s string) bool {
		return emojiVersionRegex.MatchString(s)
	})

	name := emojiFields[(emojiVersionIndex + 1):]

	return strings.Join(name, " ")
}

func ParseEmojiCharacter(emojiFields []string) string {
	emojiCharacterIndex := slices.Index(emojiFields, "#") + 1

	return emojiFields[emojiCharacterIndex]
}
