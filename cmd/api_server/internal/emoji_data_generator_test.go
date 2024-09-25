//nolint:errcheck
package data_generation

import (
	"net/http"
	"net/http/httptest"
	"testing"

	test_helpers "github.com/brandonau24/emoji-data-generator/cmd/api_server/internal/internal"
)

func Test_Generate_SkipsComments(t *testing.T) {
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == test_helpers.MOCK_UNICODE_EMOJIS_PATH {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`# This is a comment
# This is another comment
# This is the last comment`))
		}

		if r.URL.Path == test_helpers.MOCK_UNICODE_ANNOTATIONS_PATH {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
		{
			"annotations": {
				"annotations": {
					"😀": {
						"default": ["one"],
						"tts": ["grinning face"]
									}
								}
					}
		}
`))
		}
	}))
	defer mockHttpServer.Close()

	emojiDataGenerator := EmojiDataGenerator{}
	emojis, _ := emojiDataGenerator.Generate(test_helpers.MockDataUrlProvider{
		BaseUrl: mockHttpServer.URL,
	})

	if len(emojis) != 0 {
		t.Errorf("Failed to parse comments")
	}
}

func Test_Generate_SetsCodepoint(t *testing.T) {
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == test_helpers.MOCK_UNICODE_EMOJIS_PATH {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`# group: group1
1F600                                                  ; fully-qualified     # 😀 E1.0 grinning face`))
		}

		if r.URL.Path == test_helpers.MOCK_UNICODE_ANNOTATIONS_PATH {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
		{
			"annotations": {
				"annotations": {
					"😀": {
						"default": ["one"],
						"tts": ["grinning face"]
									}
								}
					}
		}
`))
		}
	}))

	defer mockHttpServer.Close()

	emojiDataGenerator := EmojiDataGenerator{}
	emojis, _ := emojiDataGenerator.Generate(test_helpers.MockDataUrlProvider{
		BaseUrl: mockHttpServer.URL,
	})
	emoji := emojis["group1"][0]

	if emoji.Codepoints != "1F600" {
		t.Errorf("Failed to parse codepoint. Received %v, expected 😀", emoji.Codepoints)
	}
}

func Test_Generate_SetsCodepoints(t *testing.T) {
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == test_helpers.MOCK_UNICODE_EMOJIS_PATH {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`# group: group1
1F62E 200D 1F4A8                                       ; fully-qualified     # 😮‍💨 E13.1 face exhaling`))
		}

		if r.URL.Path == test_helpers.MOCK_UNICODE_ANNOTATIONS_PATH {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
		{
			"annotations": {
				"annotations": {
					"😀": {
						"default": ["one"],
						"tts": ["grinning face"]
									}
								}
					}
		}
`))
		}
	}))

	defer mockHttpServer.Close()

	emojiDataGenerator := EmojiDataGenerator{}
	emojis, _ := emojiDataGenerator.Generate(test_helpers.MockDataUrlProvider{
		BaseUrl: mockHttpServer.URL,
	})
	emoji := emojis["group1"][0]

	if emoji.Codepoints != "1F62E 200D 1F4A8" {
		t.Errorf("Failed to parse codepoint. Received %v, expected 😮‍💨", emoji.Codepoints)
	}
}

func Test_Generate_SetsName_WithFirstNameFromAnnotations(t *testing.T) {
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == test_helpers.MOCK_UNICODE_EMOJIS_PATH {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`# group: group1
1F600                                                  ; fully-qualified     # 😀 E1.0 grinning face`))
		}

		if r.URL.Path == test_helpers.MOCK_UNICODE_ANNOTATIONS_PATH {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
		{
			"annotations": {
				"annotations": {
					"😀": {
						"default": ["one"],
						"tts": ["grinning face"]
									}
								}
					}
		}
`))
		}
	}))

	defer mockHttpServer.Close()

	emojiDataGenerator := EmojiDataGenerator{}
	emojis, _ := emojiDataGenerator.Generate(test_helpers.MockDataUrlProvider{
		BaseUrl: mockHttpServer.URL,
	})
	emoji := emojis["group1"][0]

	if emoji.Name != "grinning face" {
		t.Errorf("Failed to parse emoji name. Received %v, expected grinning face", emoji.Name)
	}
}

func Test_Generate_SelectsFullyQualifiedEmojis(t *testing.T) {
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == test_helpers.MOCK_UNICODE_EMOJIS_PATH {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(` # group: group1
F636 200D 1F32B FE0F                                  ; fully-qualified     # 😶‍🌫️ E13.1 face in clouds
1F636 200D 1F32B                                       ; minimally-qualified # 😶‍🌫 E13.1 face in clouds
2620                                                   ; unqualified         # ☠ E1.0 skull and crossbones
`))
		}

		if r.URL.Path == test_helpers.MOCK_UNICODE_ANNOTATIONS_PATH {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
		{
			"annotations": {
				"annotations": {
					"😀": {
						"default": ["one"],
						"tts": ["grinning face"]
									}
								}
					}
		}
`))
		}
	}))

	defer mockHttpServer.Close()

	emojiDataGenerator := EmojiDataGenerator{}
	emojis, _ := emojiDataGenerator.Generate(test_helpers.MockDataUrlProvider{
		BaseUrl: mockHttpServer.URL,
	})
	emojisInGroup1 := emojis["group1"]

	if len(emojisInGroup1) != 1 {
		t.Errorf("Other emojis that are not fully qualified have been parsed")
	}
}

func Test_Generate_GroupsEmojis(t *testing.T) {
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == test_helpers.MOCK_UNICODE_EMOJIS_PATH {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`# group: Smileys & Emotion
1F600                                                  ; fully-qualified     # 😀 E1.0 grinning face
1F603                                                  ; fully-qualified     # 😃 E0.6 grinning face with big eyes
1F636 200D 1F32B                                       ; minimally-qualified # 😶‍🌫 E13.1 face in clouds

# group: People & Body
1F44B                                                  ; fully-qualified     # 👋 E0.6 waving hand
1F44B 1F3FB                                            ; fully-qualified     # 👋🏻 E1.0 waving hand: light skin tone
1F590                                                  ; unqualified         # 🖐 E0.7 hand with fingers splayed
`))
		}

		if r.URL.Path == test_helpers.MOCK_UNICODE_ANNOTATIONS_PATH {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
		{
			"annotations": {
				"annotations": {
					"😀": {
						"default": ["one"],
						"tts": ["grinning face"]
									}
								}
					}
		}
`))
		}
	}))

	defer mockHttpServer.Close()

	emojiDataGenerator := EmojiDataGenerator{}
	emojis, _ := emojiDataGenerator.Generate(test_helpers.MockDataUrlProvider{
		BaseUrl: mockHttpServer.URL,
	})

	smileyAndEmotionsGroup, smileyAndEmotionOk := emojis["Smileys & Emotion"]

	if !smileyAndEmotionOk {
		t.Errorf("Could not parse Smileys & Emotion group")
	}

	if len(smileyAndEmotionsGroup) != 2 {
		t.Errorf("Expected 2 emojis in Smiley & Emotion group, received %v", len(smileyAndEmotionsGroup))
	}

	peopleAndBodyGroup, peopleAndBodyOk := emojis["People & Body"]

	if !peopleAndBodyOk {
		t.Errorf("Could not parse People & Body group")
	}

	if len(peopleAndBodyGroup) != 2 {
		t.Errorf("Expected 2 emojis in People & Body group, received %v", len(peopleAndBodyGroup))
	}
}

func Test_Generate_SetsAnnotations(t *testing.T) {
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == test_helpers.MOCK_UNICODE_EMOJIS_PATH {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`# group: Smileys & Emotion
1F600                                                  ; fully-qualified     # 😀 E1.0 grinning face
1F636 200D 1F32B FE0F                                  ; fully-qualified     # 😶‍🌫️ E13.1 face in clouds
1F636 200D 1F32B                                       ; minimally-qualified # 😶‍🌫 E13.1 face in clouds
`))
		}

		if r.URL.Path == test_helpers.MOCK_UNICODE_ANNOTATIONS_PATH {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
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
				},
			"😶‍🌫️": {
				"default": [
					"absentminded",
					"face in clouds",
					"face in the fog",
					"head in clouds"
				],
				"tts": ["face in clouds"]
			}
		}
	}
}
`))
		}
	}))

	defer mockHttpServer.Close()

	smileyAnnotations := []string{"face", "grin", "grinning face"}

	emojiDataGenerator := EmojiDataGenerator{}
	emojis, _ := emojiDataGenerator.Generate(test_helpers.MockDataUrlProvider{
		BaseUrl: mockHttpServer.URL,
	})

	smileyAndEmotionsGroup, ok := emojis["Smileys & Emotion"]
	if !ok {
		t.Errorf("Could not parse Smileys & Emotion group")
	}

	smileyEmoji := smileyAndEmotionsGroup[0]
	if !test_helpers.AreAnnotationsEqual(smileyEmoji.Annotations, smileyAnnotations) {
		t.Errorf("Failed to map annotations. Received %v, expected %v", smileyEmoji.Annotations, smileyAnnotations)
	}

	faceCloudAnnotations := []string{"absentminded", "face in clouds", "face in the fog", "head in clouds"}
	faceInCloudEmoji := smileyAndEmotionsGroup[1]
	if !test_helpers.AreAnnotationsEqual(faceInCloudEmoji.Annotations, faceCloudAnnotations) {
		t.Errorf("Failed to map annotations. Received %v, expected %v", faceInCloudEmoji.Annotations, faceCloudAnnotations)

	}
}

func Test_Generate_SetsCharacter(t *testing.T) {
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == test_helpers.MOCK_UNICODE_EMOJIS_PATH {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(` # group: group1
1F62E 200D 1F4A8                                       ; fully-qualified     # 😮‍💨 E13.1 face exhaling
1F44B 1F3FB                                            ; fully-qualified     # 👋🏻 E1.0 waving hand: light skin tone
`))
		}

		if r.URL.Path == test_helpers.MOCK_UNICODE_ANNOTATIONS_PATH {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
		{
			"annotations": {
				"annotations": {
					"😀": {
						"default": ["one"],
						"tts": ["grinning face"]
									}
								}
					}
		}
`))
		}
	}))

	defer mockHttpServer.Close()

	emojiDataGenerator := EmojiDataGenerator{}
	emojis, _ := emojiDataGenerator.Generate(test_helpers.MockDataUrlProvider{
		BaseUrl: mockHttpServer.URL,
	})

	faceExhalingEmoji := emojis["group1"][0]

	if faceExhalingEmoji.Character != "😮‍💨" {
		t.Errorf("Failed to get emoji character. Received %v, expected %v", faceExhalingEmoji.Character, "😮‍💨")
	}

	lightSkinWavingHandEmoji := emojis["group1"][1]

	if lightSkinWavingHandEmoji.Character != "👋🏻" {
		t.Errorf("Failed to get emoji character. Received %v, expected %v", lightSkinWavingHandEmoji.Character, "👋🏻")
	}
}

func Test_Generatate_EmojiDataFetchFails(t *testing.T) {
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == test_helpers.MOCK_UNICODE_EMOJIS_PATH {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed test request"))
		}

		if r.URL.Path == test_helpers.MOCK_UNICODE_ANNOTATIONS_PATH {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
		{
			"annotations": {
				"annotations": {
					"😀": {
						"default": ["one"],
						"tts": ["grinning face"]
									}
								}
					}
		}
`))
		}
	}))

	defer mockHttpServer.Close()

	emojiDataGenerator := EmojiDataGenerator{}
	_, err := emojiDataGenerator.Generate(test_helpers.MockDataUrlProvider{
		BaseUrl: mockHttpServer.URL,
	})

	if err == nil {
		t.Errorf("Expected parser to return error on a failed request")
	}
}

func Test_Generate_EmptyAnnotations(t *testing.T) {
	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == test_helpers.MOCK_UNICODE_EMOJIS_PATH {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed test request"))
		}

		if r.URL.Path == test_helpers.MOCK_UNICODE_ANNOTATIONS_PATH {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed test request"))
		}
	}))

	emojiDataGenerator := EmojiDataGenerator{}
	_, err := emojiDataGenerator.Generate(test_helpers.MockDataUrlProvider{
		BaseUrl: mockHttpServer.URL,
	})

	if err == nil {
		t.Errorf("Expected parser to return error on a failed request")
	}
}
