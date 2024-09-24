package providers

type DataUrlProvider interface {
	GetUnicodeEmojisDataUrl() string
	GetUnicodeAnnotationsUrl() string
}

type UnicodeDataUrlProvider struct{}

func (p UnicodeDataUrlProvider) GetUnicodeEmojisDataUrl() string {
	return "https://unicode.org/Public/emoji/15.1/emoji-test.txt"
}

func (p UnicodeDataUrlProvider) GetUnicodeAnnotationsUrl() string {
	return "https://raw.githubusercontent.com/unicode-org/cldr-json/refs/heads/main/cldr-json/cldr-annotations-modern/annotations/en/annotations.json"
}
