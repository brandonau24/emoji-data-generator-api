package parsers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/brandonau24/emoji-data-generator-api/cmd/api_server/internal/providers"
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
		log.Printf("Unicode Annotations File: %v - HTTP Status Code: %v", url, annotationsResponse.StatusCode)
		log.Printf("Unicode Annotations File: %v - Response Body: %v", url, string(responseBytes))

		return nil, fmt.Errorf("could not make successful request to %v", url)
	}

	return annotationsResponse, nil

}

func ParseAnnotations(p providers.DataUrlProvider, annotationsChannel chan map[string]Annotation, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	waitGroup.Add(1)

	annotationsResponse, err := fetchUnicodeAnnotations(p.GetUnicodeAnnotationsUrl())

	if err != nil {
		annotationsChannel <- nil

		return
	}

	defer annotationsResponse.Body.Close()

	annotationsResponseBody, readErr := io.ReadAll(annotationsResponse.Body)

	if readErr != nil {
		annotationsChannel <- nil

		return
	}

	var annotationsFileMap AnnotationsFile
	jsonErr := json.Unmarshal(annotationsResponseBody, &annotationsFileMap)

	if jsonErr != nil {
		annotationsChannel <- nil

		return
	}

	annotationsChannel <- annotationsFileMap.Annotations.Annotations
}
