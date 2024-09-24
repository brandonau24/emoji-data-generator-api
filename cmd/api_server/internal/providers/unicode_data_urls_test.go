package providers

import (
	"strings"
	"testing"
)

func TestGetUnicodeEmojisDataUrlWithoutVersion(t *testing.T) {
	urlProvider := UnicodeDataUrlProvider{}

	url := urlProvider.GetUnicodeEmojisDataUrl()

	if !strings.Contains(url, "latest") {
		t.Errorf("latest unicode version is not used for url when no version is provided. Got %v", url)

	}
}

func TestGetUnicodeEmojisDataUrlWithVersion(t *testing.T) {
	version := "15.0"
	urlProvider := UnicodeDataUrlProvider{
		Version: version,
	}

	url := urlProvider.GetUnicodeEmojisDataUrl()

	if !strings.Contains(url, version) {
		t.Errorf("%v unicode version is not used for url when provided. Got %v", version, url)

	}
}
