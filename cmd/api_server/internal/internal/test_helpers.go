package test_helpers

func AreAnnotationsEqual(annotations1 []string, annotations2 []string) bool {
	if len(annotations1) != len(annotations2) {
		return false
	}

	for i := range annotations1 {
		if annotations1[i] != annotations2[i] {
			return false
		}
	}

	return true
}

const (
	MOCK_UNICODE_ANNOTATIONS_PATH = "/mock/annotations.json"
	MOCK_UNICODE_EMOJIS_PATH      = "/mock/emoji.txt"
)

type MockDataUrlProvider struct {
	BaseUrl string
}

func (p MockDataUrlProvider) GetUnicodeEmojisDataUrl() string {
	return p.BaseUrl + MOCK_UNICODE_EMOJIS_PATH
}

func (p MockDataUrlProvider) GetUnicodeAnnotationsUrl() string {
	return p.BaseUrl + MOCK_UNICODE_ANNOTATIONS_PATH
}
