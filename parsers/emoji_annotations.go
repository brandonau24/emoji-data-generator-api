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
}

func ParseAnnotations(annotations string) map[string][]string {
	var annotationsFileMap AnnotationsFile
	err := json.Unmarshal([]byte(annotations), &annotationsFileMap)

	if err != nil {
		panic(err)
	}

	parsedAnnotations := make(map[string][]string, 0)

	nestedAnnotations := annotationsFileMap.Annotations.Annotations

	for emoji, emojiAnnotations := range nestedAnnotations {
		emojiCodepoints := []rune(emoji)
		emojiCodepointsHexadecimal := fmt.Sprintf("%X", emojiCodepoints)
		emojiCodepointsHexadecimal = strings.ReplaceAll(emojiCodepointsHexadecimal, "[", "")
		emojiCodepointsHexadecimal = strings.ReplaceAll(emojiCodepointsHexadecimal, "]", "")
		parsedAnnotations[emojiCodepointsHexadecimal] = emojiAnnotations.Default
	}

	return parsedAnnotations
}
