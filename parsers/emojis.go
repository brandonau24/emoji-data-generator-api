package parsers

import (
	"bufio"
	"regexp"
	"slices"
	"strings"
)

type Emoji struct {
	Codepoints  string   `json:"codepoints"`
	Annotations []string `json:"annotations"`
	Name        string   `json:"name"`
}

const EMOJIS_FILE_PATH = "unicode/v15.1/emojis.txt"
const FULLY_QUALIFIED = "fully-qualified"

func parseCodepoints(emojiFields []string) string {
	semicolonIndex := slices.Index(emojiFields, ";")
	codepoints := emojiFields[0:semicolonIndex]

	return strings.Join(codepoints, " ")
}

func parseEmojiName(emojiFields []string) string {
	emojiVersionRegex := regexp.MustCompile(`E\d+\.\d+`)
	emojiVersionIndex := slices.IndexFunc(emojiFields, func(s string) bool {
		return emojiVersionRegex.MatchString(s)
	})

	name := emojiFields[(emojiVersionIndex + 1):]

	return strings.Join(name, " ")
}

func ParseEmojis(e string, annotations map[string][]string) map[string][]Emoji {
	var currentGroup string

	emojis := make(map[string][]Emoji, 0)
	scanner := bufio.NewScanner(strings.NewReader(e))

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.Index(line, "#") == 0 {
			if strings.Contains(line, "# group: ") {
				startOfGroupName := strings.Index(line, ":") + 2
				currentGroup = line[startOfGroupName:]
			} else {
				continue
			}
		} else {
			emojiFields := strings.Fields(line)

			if slices.Index(emojiFields, FULLY_QUALIFIED) == -1 {
				continue
			}

			codepoints := parseCodepoints(emojiFields)
			name := parseEmojiName(emojiFields)
			emojiAnnotations := annotations[codepoints]

			newEmoji := Emoji{
				Codepoints:  codepoints,
				Name:        name,
				Annotations: emojiAnnotations,
			}

			emojisInGroup, ok := emojis[currentGroup]
			if ok {
				emojis[currentGroup] = append(emojisInGroup, newEmoji)
			} else {
				emojis[currentGroup] = []Emoji{newEmoji}
			}
		}

	}

	return emojis
}
