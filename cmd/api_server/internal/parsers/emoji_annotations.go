package parsers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AnnotationsFile struct {
	Annotations struct {
		Annotations map[string]Annotation
	}
}

type Annotation struct {
	Default []string
	Tts     []string
}

const BASE_GITHUBUSERCONTENT_URL = "https://raw.githubusercontent.com"

func fetchUnicodeAnnotations(baseUrl string) (*http.Response, error) {
	base := baseUrl
	if baseUrl == "" {
		base = BASE_GITHUBUSERCONTENT_URL
	}

	annotationsResponse, err := http.Get(fmt.Sprintf("%v/unicode-org/cldr-json/refs/heads/main/cldr-json/cldr-annotations-modern/annotations/en/annotations.json", base))

	if err != nil {
		log.Println(err.Error())

		return nil, fmt.Errorf("could not make connect to %v", baseUrl)
	}

	if annotationsResponse.StatusCode != http.StatusOK {
		responseBytes, _ := io.ReadAll(annotationsResponse.Body)
		log.Printf("Unicode Annotations File - HTTP Status Code: %v", annotationsResponse.StatusCode)
		log.Printf("Unicode Annotations File - Response Body: %v", string(responseBytes))

		return nil, fmt.Errorf("could not make successful request to %v", BASE_GITHUBUSERCONTENT_URL)
	}

	return annotationsResponse, nil

}

func ParseAnnotations(baseUrl string) map[string]Annotation {
	annotationsResponse, err := fetchUnicodeAnnotations(baseUrl)

	if err != nil {
		return nil
	}

	defer annotationsResponse.Body.Close()

	annotationsResponseBody, readErr := io.ReadAll(annotationsResponse.Body)

	if readErr != nil {
		return nil
	}

	var annotationsFileMap AnnotationsFile
	jsonErr := json.Unmarshal(annotationsResponseBody, &annotationsFileMap)

	if jsonErr != nil {
		return nil
	}

	return annotationsFileMap.Annotations.Annotations
}
