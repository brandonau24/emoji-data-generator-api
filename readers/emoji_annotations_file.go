package readers

import _ "embed"

//go:embed en.json
var annotationsFileContent string

func ReadEmojiAnnotationsFile() string {

	return annotationsFileContent
}
