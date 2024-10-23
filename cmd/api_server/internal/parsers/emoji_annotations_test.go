//nolint:errcheck
package parsers

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	test_helpers "github.com/brandonau24/emoji-data-generator-api/cmd/api_server/internal/internal"
)

type emojiAnnotationsTest struct {
	Name                string
	StatusCode          int
	ResponseBody        []byte
	ExpectedAnnotations map[string]Annotation
}

func TestParseAnnotations(t *testing.T) {
	tests := map[string]emojiAnnotationsTest{
		"maps emoji to its annotations": {
			StatusCode: http.StatusOK,
			ResponseBody: []byte(`
			{
				"annotations": {
					"annotations": {
						"😀": {
							"default": [
								"face",
								"grin",
								"grinning face"
							],
							"tts": [
								"grinning face"
							]
						}
					}
				}
			}`),
			ExpectedAnnotations: map[string]Annotation{
				"😀": {
					Tts:     []string{"grinning face"},
					Default: []string{"face", "grin", "grinning face"},
				},
			},
		},
		"nil is returned when request fails": {
			StatusCode:          http.StatusBadRequest,
			ResponseBody:        []byte(""),
			ExpectedAnnotations: nil,
		},
	}

	for name, testcase := range tests {
		t.Run(name, func(t *testing.T) {
			mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == test_helpers.MOCK_UNICODE_ANNOTATIONS_PATH {
					w.WriteHeader(testcase.StatusCode)
					w.Write(testcase.ResponseBody)
				}
			}))

			defer mockHttpServer.Close()

			annotations := ParseAnnotations(test_helpers.MockDataUrlProvider{
				BaseUrl: mockHttpServer.URL,
			})

			if !reflect.DeepEqual(annotations, testcase.ExpectedAnnotations) {
				t.Errorf("%v: expected: %v, got: %v", testcase.Name, testcase.ExpectedAnnotations, annotations)
			}
		})
	}
}
