package parsers

import (
	"encoding/json"
	"fmt"
)

type AnnotationsFile struct {
	Annotations struct {
		Annotations map[string]Annotation
	}
}

type Annotation struct {
	Default []string
}

type EmojiHexadecimalCodepoint string

func (codepoint EmojiHexadecimalCodepoint) MarshalJSON() ([]byte, error) {
	codepointStr := fmt.Sprintf("\"%s\"", string(codepoint))

	return json.Marshal(codepointStr)
}

func ParseAnnotations(annotations string) map[EmojiHexadecimalCodepoint][]string {
	var annotationsFileMap AnnotationsFile
	err := json.Unmarshal([]byte(annotations), &annotationsFileMap)

	if err != nil {
		panic(err)
	}

	parsedAnnotations := make(map[EmojiHexadecimalCodepoint][]string, 0)

	nestedAnnotations := annotationsFileMap.Annotations.Annotations

	for emoji, emojiAnnotations := range nestedAnnotations {
		emojiCodepoints := []rune(emoji)
		emojiCodepointsHexadecimal := fmt.Sprintf("%X", emojiCodepoints)
		parsedAnnotations[EmojiHexadecimalCodepoint(emojiCodepointsHexadecimal)] = emojiAnnotations.Default
	}

	return parsedAnnotations
}
