package parsers

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

type Emoji struct {
	codepoints  string
	annotations []string
	name        string
}

const EMOJIS_FILE_PATH = "unicode/v15.1/emojis.txt"

func ParseEmojis() []Emoji {
	emojis := make([]Emoji, 3773) // Number of fully-qualified emojis

	workingDir, _ := os.Getwd()
	filePath := path.Join(workingDir, EMOJIS_FILE_PATH)
	emojiFile, err := os.Open(filePath)

	if err != nil {
		panic(fmt.Sprintf("cannot open file: %v", filePath))
	}

	scanner := bufio.NewScanner(emojiFile)

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}

	return emojis
}
