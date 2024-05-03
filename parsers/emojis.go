package parsers

import (
	"bufio"
	"regexp"
	"slices"
	"strings"
)

type Emoji struct {
	Codepoints  string
	annotations []string
	Name        string
}

const EMOJIS_FILE_PATH = "unicode/v15.1/emojis.txt"
const FULLY_QUALIFIED = "fully-qualified"

func ParseEmojis(e string) []Emoji {
	emojis := make([]Emoji, 0)
	scanner := bufio.NewScanner(strings.NewReader(e))

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.Index(line, "#") == 0 {
			continue
		} else {
			emojiFields := strings.Fields(line)

			if slices.Index(emojiFields, FULLY_QUALIFIED) == -1 {
				continue
			}

			semicolonIndex := slices.Index(emojiFields, ";")
			codepoints := emojiFields[0:semicolonIndex]

			emojiVersionRegex := regexp.MustCompile(`E\d+\.\d+`)
			emojiVersionIndex := slices.IndexFunc(emojiFields, func(s string) bool {
				return emojiVersionRegex.MatchString(s)
			})

			name := emojiFields[(emojiVersionIndex + 1):]

			emojis = append(emojis, Emoji{
				Codepoints: strings.Join(codepoints, " "),
				Name:       strings.Join(name, " "),
			})
		}

	}

	return emojis
}
