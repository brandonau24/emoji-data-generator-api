package parsers

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"slices"
	"strings"
)

type Emoji struct {
	Character   string   `json:"character"`
	Codepoints  string   `json:"codepoints"`
	Annotations []string `json:"annotations"`
	Name        string   `json:"name"`
}

const (
	FULLY_QUALIFIED  = "fully-qualified"
	UNICODE_BASE_URL = "https://unicode.org"
)

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

func parseEmojiCharacter(emojiFields []string) string {
	emojiCharacterIndex := slices.Index(emojiFields, "#") + 1

	return emojiFields[emojiCharacterIndex]
}

func fetchEmojiDataFile(unicodeBaseUrl string) (*http.Response, error) {
	baseUrl := unicodeBaseUrl
	if baseUrl == "" {
		baseUrl = UNICODE_BASE_URL
	}

	emojisResponse, err := http.Get(fmt.Sprintf("%v/Public/emoji/16.0/emoji-test.txt", baseUrl))

	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("could not make connect to %v", baseUrl)
	} else if emojisResponse.StatusCode != http.StatusOK {
		responseBytes, _ := io.ReadAll(emojisResponse.Body)
		log.Printf("Emojis Data File - HTTP Status Code: %v", emojisResponse.StatusCode)
		log.Printf("Emojis Data File - Response Body: %v", string(responseBytes))

		return nil, fmt.Errorf("could not make successful request to unicode.org")
	}

	return emojisResponse, nil
}

func ParseEmojis(unicodeBaseUrl string, annotations map[string]Annotation) (map[string][]Emoji, error) {
	var currentGroup string

	emojis := make(map[string][]Emoji, 0)

	emojiDataFileResponse, fetchErr := fetchEmojiDataFile(unicodeBaseUrl)

	if fetchErr != nil {
		return nil, fetchErr
	}

	defer emojiDataFileResponse.Body.Close()

	scanner := bufio.NewScanner(emojiDataFileResponse.Body)

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

			character := parseEmojiCharacter(emojiFields)
			codepoints := parseCodepoints(emojiFields)
			name := parseEmojiName(emojiFields)
			emojiAnnotations := annotations[codepoints]

			newEmoji := Emoji{
				Character:   character,
				Codepoints:  codepoints,
				Name:        name,
				Annotations: emojiAnnotations.Default,
			}

			emojisInGroup, ok := emojis[currentGroup]
			if ok {
				emojis[currentGroup] = append(emojisInGroup, newEmoji)
			} else {
				emojis[currentGroup] = []Emoji{newEmoji}
			}
		}

	}

	return emojis, nil
}
