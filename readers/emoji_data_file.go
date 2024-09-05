package readers

import _ "embed"

//go:embed emojis.txt
var emojiDataFileContent string

func ReadEmojiDataFile() string {
	return emojiDataFileContent
}
