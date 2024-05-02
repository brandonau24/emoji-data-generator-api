package readers

import (
	"fmt"
	"os"
	"path"
)

const EMOJIS_FILE_PATH = "unicode/v15.1/emojis.txt"

func ReadEmojiDataFile() string {
	workingDir, _ := os.Getwd()
	filePath := path.Join(workingDir, EMOJIS_FILE_PATH)
	emojiFile, err := os.ReadFile(filePath)

	if err != nil {
		panic(fmt.Sprintf("cannot open file: %v", filePath))
	}

	return string(emojiFile)
}
