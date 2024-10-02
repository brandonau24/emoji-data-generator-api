//nolint:errcheck
package data_generation

import (
	"net/http"
	"net/http/httptest"
	"testing"

	test_helpers "github.com/brandonau24/emoji-data-generator/cmd/api_server/internal/internal"
)

const version = 1.0

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

	emojiDataGenerator := EmojiDataGenerator{
		UrlProvider: test_helpers.MockDataUrlProvider{
			BaseUrl: mockHttpServer.URL,
		},
	}
	emojis, _ := emojiDataGenerator.Generate(1.0)

	if len(emojis) != 0 {
		t.Errorf("Failed to parse comments")
	}
}

func Test_Generate_SetsCodepoint(t *testing.T) {
	mockDataUrlProvider := test_helpers.MockDataUrlProvider{}
	mockEmojiPath := mockDataUrlProvider.BuildUrlPath(version)

	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == mockEmojiPath {
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

	mockDataUrlProvider.BaseUrl = mockHttpServer.URL

	emojiDataGenerator := EmojiDataGenerator{
		UrlProvider: test_helpers.MockDataUrlProvider{
			BaseUrl: mockHttpServer.URL,
		},
	}
	emojis, _ := emojiDataGenerator.Generate(version)
	emoji := emojis["group1"][0]

	if emoji.Codepoints != "1F600" {
		t.Errorf("Failed to parse codepoint. Received %v, expected 😀", emoji.Codepoints)
	}
}

func Test_Generate_SetsCodepoints(t *testing.T) {
	mockDataUrlProvider := test_helpers.MockDataUrlProvider{}
	mockEmojiPath := mockDataUrlProvider.BuildUrlPath(version)

	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == mockEmojiPath {
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

	mockDataUrlProvider.BaseUrl = mockHttpServer.URL

	emojiDataGenerator := EmojiDataGenerator{
		UrlProvider: mockDataUrlProvider,
	}
	emojis, _ := emojiDataGenerator.Generate(version)
	emoji := emojis["group1"][0]

	if emoji.Codepoints != "1F62E 200D 1F4A8" {
		t.Errorf("Failed to parse codepoint. Received %v, expected 😮‍💨", emoji.Codepoints)
	}
}

func Test_Generate_SetsName_WithFirstNameFromAnnotations(t *testing.T) {
	mockDataUrlProvider := test_helpers.MockDataUrlProvider{}
	mockEmojiPath := mockDataUrlProvider.BuildUrlPath(version)

	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == mockEmojiPath {
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

	mockDataUrlProvider.BaseUrl = mockHttpServer.URL

	emojiDataGenerator := EmojiDataGenerator{
		UrlProvider: mockDataUrlProvider,
	}
	emojis, _ := emojiDataGenerator.Generate(version)
	emoji := emojis["group1"][0]

	if emoji.Name != "grinning face" {
		t.Errorf("Failed to parse emoji name. Received %v, expected grinning face", emoji.Name)
	}
}

func Test_Generate_SetsName_FromEmojiData_WhenAnnotationsAreMissing(t *testing.T) {
	mockDataUrlProvider := test_helpers.MockDataUrlProvider{}
	mockEmojiPath := mockDataUrlProvider.BuildUrlPath(version)

	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == mockEmojiPath {
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
					"😝": {
						"default": ["one"],
						"tts": ["tongue face"]
									}
								}
					}
		}
`))
		}
	}))

	defer mockHttpServer.Close()

	mockDataUrlProvider.BaseUrl = mockHttpServer.URL

	emojiDataGenerator := EmojiDataGenerator{
		UrlProvider: mockDataUrlProvider,
	}
	emojis, _ := emojiDataGenerator.Generate(version)
	emoji := emojis["group1"][0]

	if emoji.Name != "grinning face" {
		t.Errorf("Failed to parse emoji name. Received %v, expected grinning face", emoji.Name)
	}
}

func Test_Generate_SelectsFullyQualifiedEmojis(t *testing.T) {
	mockDataUrlProvider := test_helpers.MockDataUrlProvider{}
	mockEmojiPath := mockDataUrlProvider.BuildUrlPath(version)

	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == mockEmojiPath {
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

	mockDataUrlProvider.BaseUrl = mockHttpServer.URL

	emojiDataGenerator := EmojiDataGenerator{
		UrlProvider: mockDataUrlProvider,
	}
	emojis, _ := emojiDataGenerator.Generate(version)
	emojisInGroup1 := emojis["group1"]

	if len(emojisInGroup1) != 1 {
		t.Errorf("Other emojis that are not fully qualified have been parsed")
	}
}

func Test_Generate_GroupsEmojis(t *testing.T) {
	mockDataUrlProvider := test_helpers.MockDataUrlProvider{}
	mockEmojiPath := mockDataUrlProvider.BuildUrlPath(version)

	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == mockEmojiPath {
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

	mockDataUrlProvider.BaseUrl = mockHttpServer.URL

	emojiDataGenerator := EmojiDataGenerator{
		UrlProvider: mockDataUrlProvider,
	}
	emojis, _ := emojiDataGenerator.Generate(version)

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
	mockDataUrlProvider := test_helpers.MockDataUrlProvider{}
	mockEmojiPath := mockDataUrlProvider.BuildUrlPath(version)

	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == mockEmojiPath {
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

	mockDataUrlProvider.BaseUrl = mockHttpServer.URL

	smileyAnnotations := []string{"face", "grin", "grinning face"}

	emojiDataGenerator := EmojiDataGenerator{
		UrlProvider: mockDataUrlProvider,
	}
	emojis, _ := emojiDataGenerator.Generate(version)

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

func Test_Generate_SetsAnnotationsToNull_WhenItDoesNotExist(t *testing.T) {
	mockDataUrlProvider := test_helpers.MockDataUrlProvider{}
	mockEmojiPath := mockDataUrlProvider.BuildUrlPath(version)

	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == mockEmojiPath {
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
				}
		}
	}
}
`))
		}
	}))

	defer mockHttpServer.Close()

	mockDataUrlProvider.BaseUrl = mockHttpServer.URL

	emojiDataGenerator := EmojiDataGenerator{
		UrlProvider: mockDataUrlProvider,
	}
	emojis, _ := emojiDataGenerator.Generate(version)

	smileyAndEmotionsGroup, ok := emojis["Smileys & Emotion"]
	if !ok {
		t.Errorf("Could not parse Smileys & Emotion group")
	}

	faceInCloudEmojiAnnotations := smileyAndEmotionsGroup[1].Annotations
	if faceInCloudEmojiAnnotations != nil {
		t.Errorf("Annotations do not exist, but received non-null value: %v", faceInCloudEmojiAnnotations)
	}
}

func Test_Generate_SetsCharacter(t *testing.T) {
	mockDataUrlProvider := test_helpers.MockDataUrlProvider{}
	mockEmojiPath := mockDataUrlProvider.BuildUrlPath(version)

	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == mockEmojiPath {
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

	mockDataUrlProvider.BaseUrl = mockHttpServer.URL

	emojiDataGenerator := EmojiDataGenerator{
		UrlProvider: mockDataUrlProvider,
	}
	emojis, _ := emojiDataGenerator.Generate(version)

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
	mockDataUrlProvider := test_helpers.MockDataUrlProvider{}
	mockEmojiPath := mockDataUrlProvider.BuildUrlPath(version)

	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == mockEmojiPath {
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

	mockDataUrlProvider.BaseUrl = mockHttpServer.URL

	emojiDataGenerator := EmojiDataGenerator{
		UrlProvider: mockDataUrlProvider,
	}
	_, err := emojiDataGenerator.Generate(version)

	if err == nil {
		t.Errorf("Expected parser to return error on a failed request")
	}
}

func Test_Generate_EmojiDataResponseIsEmpty(t *testing.T) {
	mockDataUrlProvider := test_helpers.MockDataUrlProvider{}
	mockEmojiPath := mockDataUrlProvider.BuildUrlPath(version)

	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == mockEmojiPath {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(""))
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

	mockDataUrlProvider.BaseUrl = mockHttpServer.URL

	emojiDataGenerator := EmojiDataGenerator{
		UrlProvider: mockDataUrlProvider,
	}
	_, err := emojiDataGenerator.Generate(version)

	if err == nil {
		t.Errorf("Expected generator to return error on an empty request")
	}
}

func Test_Generate_EmptyAnnotations(t *testing.T) {
	mockDataUrlProvider := test_helpers.MockDataUrlProvider{}
	mockEmojiPath := mockDataUrlProvider.BuildUrlPath(version)

	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == mockEmojiPath {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed test request"))
		}

		if r.URL.Path == test_helpers.MOCK_UNICODE_ANNOTATIONS_PATH {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed test request"))
		}
	}))

	defer mockHttpServer.Close()

	mockDataUrlProvider.BaseUrl = mockHttpServer.URL

	emojiDataGenerator := EmojiDataGenerator{
		UrlProvider: mockDataUrlProvider,
	}
	_, err := emojiDataGenerator.Generate(version)

	if err == nil {
		t.Errorf("Expected parser to return error on a failed request")
	}
}

func Test_Generate_MultipleUnicodeVersions(t *testing.T) {
	version1 := 5.0
	version2 := 10.0

	mockDataUrlProvider := test_helpers.MockDataUrlProvider{}
	mockEmojiPathVersion1 := mockDataUrlProvider.BuildUrlPath(version1)
	mockEmojiPathVersion2 := mockDataUrlProvider.BuildUrlPath(version2)

	mockHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == mockEmojiPathVersion1 {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`# group: group1
1F600                                                  ; fully-qualified     # 😀 E1.0 grinning face
1F603                                                  ; fully-qualified     # 😃 E0.6 grinning face with big eyes
1F604                                                  ; fully-qualified     # 😄 E0.6 grinning face with smiling eyes
1F601                                                  ; fully-qualified     # 😁 E0.6 beaming face with smiling eyes
1F606                                                  ; fully-qualified     # 😆 E0.6 grinning squinting face
1F605                                                  ; fully-qualified     # 😅 E0.6 grinning face with sweat
1F923                                                  ; fully-qualified     # 🤣 E3.0 rolling on the floor laughing
`))
		}

		if r.URL.Path == mockEmojiPathVersion2 {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`# group: group1
1F600                                                  ; fully-qualified     # 😀 E1.0 grinning face
1F603                                                  ; fully-qualified     # 😃 E0.6 grinning face with big eyes
1F604                                                  ; fully-qualified     # 😄 E0.6 grinning face with smiling eyes
1F601                                                  ; fully-qualified     # 😁 E0.6 beaming face with smiling eyes
1F606                                                  ; fully-qualified     # 😆 E0.6 grinning squinting face
1F605                                                  ; fully-qualified     # 😅 E0.6 grinning face with sweat
1F923                                                  ; fully-qualified     # 🤣 E3.0 rolling on the floor laughing
1F602                                                  ; fully-qualified     # 😂 E0.6 face with tears of joy
1F642                                                  ; fully-qualified     # 🙂 E1.0 slightly smiling face
1F643                                                  ; fully-qualified     # 🙃 E1.0 upside-down face
1FAE0                                                  ; fully-qualified     # 🫠 E14.0 melting face
1F609                                                  ; fully-qualified     # 😉 E0.6 winking face
1F60A                                                  ; fully-qualified     # 😊 E0.6 smiling face with smiling eyes
1F607                                                  ; fully-qualified     # 😇 E1.0 smiling face with halo
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

	mockDataUrlProvider.BaseUrl = mockHttpServer.URL

	emojiDataGenerator := EmojiDataGenerator{
		UrlProvider: mockDataUrlProvider,
	}

	version1Emojis, _ := emojiDataGenerator.Generate(version1)
	version2Emojis, _ := emojiDataGenerator.Generate(version2)

	if len(version2Emojis) < len(version1Emojis) {
		t.Errorf(`
		expected %v emojis to have more than %v emojis
		Version 1 Emojis: %v
		Version 2 Emojis: %v
`, version, version2, version1Emojis, version2Emojis)
	}
}
