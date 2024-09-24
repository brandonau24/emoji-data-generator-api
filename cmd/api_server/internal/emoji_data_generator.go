package data_generation

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/brandonau24/emoji-data-generator/cmd/api_server/internal/parsers"
	"github.com/brandonau24/emoji-data-generator/cmd/api_server/internal/providers"
)

const (
	FULLY_QUALIFIED = "fully-qualified"
)

type EmojiDataGenerator struct{}

type AnnotationsFile struct {
	Annotations struct {
		Annotations map[string]Annotation
	}
}

type Annotation struct {
	Default []string
	Tts     []string
}

type Emoji struct {
	Character   string   `json:"character"`
	Codepoints  string   `json:"codepoints"`
	Annotations []string `json:"annotations"`
	Name        string   `json:"name"`
}

func fetchEmojiDataFile(url string) (*http.Response, error) {
	emojisResponse, err := http.Get(url)

	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("could not make connect to %v", url)
	} else if emojisResponse.StatusCode != http.StatusOK {
		responseBytes, _ := io.ReadAll(emojisResponse.Body)
		log.Printf("Emojis Data File - HTTP Status Code: %v", emojisResponse.StatusCode)
		log.Printf("Emojis Data File - Response Body: %v", string(responseBytes))

		return nil, fmt.Errorf("could not make successful request to %v", url)
	}

	return emojisResponse, nil
}

func (g EmojiDataGenerator) Generate(urlProvider providers.DataUrlProvider) (map[string][]Emoji, error) {
	annotations := parsers.ParseAnnotations(urlProvider)

	var currentGroup string

	emojis := make(map[string][]Emoji, 0)

	emojiDataFileResponse, fetchErr := fetchEmojiDataFile(urlProvider.GetUnicodeEmojisDataUrl())

	if fetchErr != nil {
		return nil, fetchErr
	}

	defer emojiDataFileResponse.Body.Close()

	scanner := bufio.NewScanner(emojiDataFileResponse.Body)

	emojiParser := parsers.EmojiParser{}

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

			character := emojiParser.ParseEmojiCharacter(emojiFields)
			codepoints := emojiParser.ParseCodepoints(emojiFields)
			emojiAnnotations := annotations[character]

			var name string
			if len(emojiAnnotations.Tts) > 0 {
				name = emojiAnnotations.Tts[0]
			} else {
				name = emojiParser.ParseEmojiName(emojiFields)
			}

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
