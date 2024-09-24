package providers

import "fmt"

type DataUrlProvider interface {
	GetUnicodeEmojisDataUrl() string
	GetUnicodeAnnotationsUrl() string
}

type UnicodeDataUrlProvider struct {
	Version string
}

func (p UnicodeDataUrlProvider) GetUnicodeEmojisDataUrl() string {
	url := "https://unicode.org/Public/emoji/%v/emoji-test.txt"

	if p.Version == "" {
		return fmt.Sprintf(url, "latest")
	}

	return fmt.Sprintf(url, p.Version)
}

func (p UnicodeDataUrlProvider) GetUnicodeAnnotationsUrl() string {
	return "https://raw.githubusercontent.com/unicode-org/cldr-json/refs/heads/main/cldr-json/cldr-annotations-modern/annotations/en/annotations.json"
}
