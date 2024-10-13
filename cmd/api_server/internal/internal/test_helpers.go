package test_helpers

import "fmt"

const (
	MOCK_UNICODE_ANNOTATIONS_PATH = "/mock/annotations.json"
	MOCK_UNICODE_EMOJIS_PATH      = "/mock/%v/emoji.txt"
)

type MockDataUrlProvider struct {
	BaseUrl string
}

func (p MockDataUrlProvider) BuildUrlPath(version float64) string {
	return fmt.Sprintf(MOCK_UNICODE_EMOJIS_PATH, version)
}

func (p MockDataUrlProvider) GetUnicodeEmojisDataUrl(version float64) string {
	return p.BaseUrl + p.BuildUrlPath(version)
}

func (p MockDataUrlProvider) GetUnicodeAnnotationsUrl() string {
	return p.BaseUrl + MOCK_UNICODE_ANNOTATIONS_PATH
}
