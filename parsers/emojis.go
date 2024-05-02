package parsers

import (
	"bufio"
	"fmt"
	"strings"
)

type Emoji struct {
	codepoints  string
	annotations []string
	name        string
}

const EMOJIS_FILE_PATH = "unicode/v15.1/emojis.txt"

func ParseEmojis(e string) []Emoji {
	emojis := make([]Emoji, 3773) // Number of fully-qualified emojis

	scanner := bufio.NewScanner(strings.NewReader(e))

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}

	return emojis
}
