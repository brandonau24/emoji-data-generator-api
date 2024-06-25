package main

import (
	"encoding/json"
	"fmt"

	"github.com/brandonau24/emoji-data-generator/parsers"
	"github.com/brandonau24/emoji-data-generator/readers"
)

func main() {
	emojiDataFile := readers.ReadEmojiDataFile()
	emojiAnnotationsFile := readers.ReadEmojiAnnotationsFile()

	emojis := parsers.ParseEmojis(emojiDataFile)
	emojisJson, err := json.Marshal(emojis)

	if err == nil {
		fmt.Println(string(emojisJson))
	}

	fmt.Println(emojiAnnotationsFile)
}
