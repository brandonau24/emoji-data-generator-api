package providers

import "fmt"

type DataUrlProvider interface {
	GetUnicodeEmojisDataUrl(version float64) string
	GetUnicodeAnnotationsUrl() string
}

type UnicodeDataUrlProvider struct{}

func (p UnicodeDataUrlProvider) GetUnicodeEmojisDataUrl(version float64) string {
	formattedVersion := fmt.Sprintf("%.1f", version)

	url := "https://unicode.org/Public/emoji/%v/emoji-test.txt"

	if version == 0.0 {
		return fmt.Sprintf(url, "latest")
	}

	return fmt.Sprintf(url, formattedVersion)
}

func (p UnicodeDataUrlProvider) GetUnicodeAnnotationsUrl() string {
	return "https://raw.githubusercontent.com/unicode-org/cldr-json/refs/heads/main/cldr-json/cldr-annotations-full/annotations/en/annotations.json"
}
