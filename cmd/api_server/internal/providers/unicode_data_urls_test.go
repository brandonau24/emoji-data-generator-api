package providers

import (
	"strings"
	"testing"
)

func TestGetUnicodeEmojisDataUrl(t *testing.T) {
	tests := map[string]struct {
		version                float64
		expectedUnicodeVersion string
	}{
		"sets unicode version to latest when input version is 0": {
			version:                0,
			expectedUnicodeVersion: "latest",
		},
		"sets unicode version based on input": {
			version:                15.0,
			expectedUnicodeVersion: "15.0",
		},
		"truncates input version to one decimal place": {
			version:                15.11,
			expectedUnicodeVersion: "15.1",
		},
		"rounds input version to one decimal place": {
			version:                15.15,
			expectedUnicodeVersion: "15.2",
		},
		"fixes input version 1 to one decimal place": {
			version:                1,
			expectedUnicodeVersion: "1.0",
		},
		"fixes input version 10 to one decimal place": {
			version:                10,
			expectedUnicodeVersion: "10.0",
		},
	}

	for name, testcase := range tests {
		t.Run(name, func(t *testing.T) {
			urlProvider := UnicodeDataUrlProvider{}

			url := urlProvider.GetUnicodeEmojisDataUrl(testcase.version)

			if !strings.Contains(url, testcase.expectedUnicodeVersion) {
				t.Errorf("%v: expected url to contain: %v, got: %v", name, testcase.expectedUnicodeVersion, url)
			}
		})
	}
}
