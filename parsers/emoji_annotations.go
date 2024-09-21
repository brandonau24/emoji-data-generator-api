package parsers

import (
	"encoding/json"
	"fmt"
	"strings"
)

type AnnotationsFile struct {
	Annotations struct {
		Annotations map[string]Annotation
	}
}

type Annotation struct {
	Default []string
	Tts     []string
}

func ParseAnnotations(annotations string) map[string]Annotation {
	var annotationsFileMap AnnotationsFile
	err := json.Unmarshal([]byte(annotations), &annotationsFileMap)

	if err != nil {
		panic(err)
	}

	parsedAnnotations := make(map[string]Annotation, 0)

	nestedAnnotations := annotationsFileMap.Annotations.Annotations

	for emoji, emojiAnnotations := range nestedAnnotations {
		emojiCodepoints := []rune(emoji)
		emojiCodepointsHexadecimal := fmt.Sprintf("%X", emojiCodepoints)
		emojiCodepointsHexadecimal = strings.ReplaceAll(emojiCodepointsHexadecimal, "[", "")
		emojiCodepointsHexadecimal = strings.ReplaceAll(emojiCodepointsHexadecimal, "]", "")
		parsedAnnotations[emojiCodepointsHexadecimal] = emojiAnnotations
	}

	return parsedAnnotations
}
