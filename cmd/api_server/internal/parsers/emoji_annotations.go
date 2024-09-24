package parsers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/brandonau24/emoji-data-generator/cmd/api_server/internal/providers"
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

func fetchUnicodeAnnotations(url string) (*http.Response, error) {
	annotationsResponse, err := http.Get(url)

	if err != nil {
		log.Println(err.Error())

		return nil, fmt.Errorf("could not make connect to %v", url)
	}

	if annotationsResponse.StatusCode != http.StatusOK {
		responseBytes, _ := io.ReadAll(annotationsResponse.Body)
		log.Printf("Unicode Annotations File - HTTP Status Code: %v", annotationsResponse.StatusCode)
		log.Printf("Unicode Annotations File - Response Body: %v", string(responseBytes))

		return nil, fmt.Errorf("could not make successful request to %v", url)
	}

	return annotationsResponse, nil

}

func ParseAnnotations(p providers.DataUrlProvider) map[string]Annotation {
	annotationsResponse, err := fetchUnicodeAnnotations(p.GetUnicodeAnnotationsUrl())

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
