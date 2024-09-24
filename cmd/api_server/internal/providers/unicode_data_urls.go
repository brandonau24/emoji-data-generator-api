package providers

import "fmt"

type DataUrlProvider interface {
	GetUnicodeEmojisDataUrl() string
	GetUnicodeAnnotationsUrl() string
}

type UnicodeDataUrlProvider struct {
	Version float64
}

func (p UnicodeDataUrlProvider) GetUnicodeEmojisDataUrl() string {
	formattedVersion := fmt.Sprintf("%.1f", p.Version)

	url := "https://unicode.org/Public/emoji/%v/emoji-test.txt"

	if p.Version == 0.0 {
		return fmt.Sprintf(url, "latest")
	}

	return fmt.Sprintf(url, formattedVersion)
}

func (p UnicodeDataUrlProvider) GetUnicodeAnnotationsUrl() string {
	return "https://raw.githubusercontent.com/unicode-org/cldr-json/refs/heads/main/cldr-json/cldr-annotations-modern/annotations/en/annotations.json"
}
