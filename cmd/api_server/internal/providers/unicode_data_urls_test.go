package providers

import (
	"strings"
	"testing"
)

func Test_GetUnicodeEmojisDataUrl_WithoutVersion_ReturnsLatest(t *testing.T) {
	urlProvider := UnicodeDataUrlProvider{}

	url := urlProvider.GetUnicodeEmojisDataUrl()

	if !strings.Contains(url, "latest") {
		t.Errorf("latest unicode version is not used for url when no version is provided. Got %v", url)

	}
}

func Test_GetUnicodeEmojisDataUrl_WithVersion(t *testing.T) {
	version := "15.0"
	urlProvider := UnicodeDataUrlProvider{
		Version: 15.0,
	}

	url := urlProvider.GetUnicodeEmojisDataUrl()

	if !strings.Contains(url, version) {
		t.Errorf("%v unicode version is not used for url when provided. Got %v", version, url)

	}
}

func Test_GetUnicodeEmojisDataUrl_TruncatesVersionToOneDecimalPlace(t *testing.T) {
	var version float64
	version = 15.11

	var urlProvider DataUrlProvider
	urlProvider = UnicodeDataUrlProvider{
		Version: version,
	}

	var url string
	url = urlProvider.GetUnicodeEmojisDataUrl()

	if !strings.Contains(url, "/15.1/") {
		t.Errorf("%v unicode version was not shortened. Got %v", version, url)

	}

	version = 15.15

	urlProvider = UnicodeDataUrlProvider{
		Version: version,
	}

	url = urlProvider.GetUnicodeEmojisDataUrl()

	if !strings.Contains(url, "/15.2/") {
		t.Errorf("%v unicode version was not shortened. Got %v", version, url)

	}

	version = 1
	urlProvider = UnicodeDataUrlProvider{
		Version: version,
	}

	url = urlProvider.GetUnicodeEmojisDataUrl()

	if !strings.Contains(url, "/1.0/") {
		t.Errorf("%v unicode version was not shortened. Got %v", version, url)

	}
}
