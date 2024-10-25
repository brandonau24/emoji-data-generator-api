package data_generation

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"strings"
	"sync"

	"github.com/brandonau24/emoji-data-generator-api/cmd/api_server/internal/parsers"
	"github.com/brandonau24/emoji-data-generator-api/cmd/api_server/internal/providers"
)

const (
	FULLY_QUALIFIED = "fully-qualified"
)

type Generator interface {
	Generate() map[string][]Emoji
}

type EmojiDataGenerator struct {
	UrlProvider providers.DataUrlProvider
}

type AnnotationsFile struct {
	Annotations struct {
		Annotations map[string]parsers.Annotation
	}
}

type Emoji struct {
	Character   string   `json:"character"`
	Codepoints  string   `json:"codepoints"`
	Annotations []string `json:"annotations"`
	Name        string   `json:"name"`
}

type emojiDataFileResponseChannel struct {
	response *http.Response
	fetchErr error
}

func fetchEmojiDataFile(url string, channel chan emojiDataFileResponseChannel, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	waitGroup.Add(1)

	emojisResponse, err := http.Get(url)

	if err != nil {
		log.Println(err.Error())
		channel <- emojiDataFileResponseChannel{response: nil, fetchErr: fmt.Errorf("could not make connect to %v", url)}
	} else if emojisResponse.StatusCode != http.StatusOK {
		responseBytes, _ := io.ReadAll(emojisResponse.Body)
		log.Printf("Emojis Data File - HTTP Status Code: %v", emojisResponse.StatusCode)
		log.Printf("Emojis Data File - Response Body: %v", string(responseBytes))

		channel <- emojiDataFileResponseChannel{response: nil, fetchErr: fmt.Errorf("could not make successful request to %v", url)}
	}

	channel <- emojiDataFileResponseChannel{response: emojisResponse, fetchErr: nil}
}

func (g EmojiDataGenerator) Generate(version float64) (map[string][]Emoji, error) {
	var waitGroup sync.WaitGroup
	var annotationsChannel = make(chan map[string]parsers.Annotation)
	var emojiDataFileResponseChannel = make(chan emojiDataFileResponseChannel)

	go parsers.ParseAnnotations(g.UrlProvider, annotationsChannel, &waitGroup)
	go fetchEmojiDataFile(g.UrlProvider.GetUnicodeEmojisDataUrl(version), emojiDataFileResponseChannel, &waitGroup)

	waitGroup.Wait()

	annotations := <-annotationsChannel
	emojiDataFileResponse := <-emojiDataFileResponseChannel

	if len(annotations) == 0 || emojiDataFileResponse.fetchErr != nil {
		return nil, fmt.Errorf("could not get unicode data")
	}

	defer emojiDataFileResponse.response.Body.Close()
	var currentGroup string

	emojis := make(map[string][]Emoji, 0)

	scanner := bufio.NewScanner(emojiDataFileResponse.response.Body)

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

			character := parsers.ParseEmojiCharacter(emojiFields)
			codepoints := parsers.ParseCodepoints(emojiFields)
			emojiAnnotations := annotations[character]

			var name string
			if len(emojiAnnotations.Tts) > 0 {
				name = emojiAnnotations.Tts[0]
			} else {
				name = parsers.ParseEmojiName(emojiFields)
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

	if len(emojis) == 0 {
		return nil, fmt.Errorf("could not generate emoji data")
	}

	return emojis, nil
}
