package readers

import (
	"fmt"
	"os"
	"path"
)

const EMOJIS_ANNOTATIONS_FILE_PATH = "unicode/cldr/v45/annotations/en.json"

func ReadEmojiAnnotationsFile() string {
	workingDir, _ := os.Getwd()
	filePath := path.Join(workingDir, EMOJIS_ANNOTATIONS_FILE_PATH)
	emojiAnnotationsFile, err := os.ReadFile(filePath)

	if err != nil {
		panic(fmt.Sprintf("cannot open file: %v", filePath))
	}

	return string(emojiAnnotationsFile)
}
