package parsers

import (
	"bufio"
	"strings"
)

type Emoji struct {
	Codepoints  string
	annotations []string
	name        string
}

const EMOJIS_FILE_PATH = "unicode/v15.1/emojis.txt"

func ParseEmojis(e string) []Emoji {
	emojis := make([]Emoji, 0)
	scanner := bufio.NewScanner(strings.NewReader(e))

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.Index(line, "#") == 0 {
			continue
		}
	}

	return emojis
}
