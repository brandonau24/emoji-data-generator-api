package main

import (
	"github.com/brandonau24/emoji-data-generator/parsers"
	"github.com/brandonau24/emoji-data-generator/readers"
)

func main() {
	emojiDataFile := readers.ReadEmojiDataFile()

	parsers.ParseEmojis(emojiDataFile)
}
