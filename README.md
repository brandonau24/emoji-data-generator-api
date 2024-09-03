# emoji-data-generator
emoji-data-generator is a small GET API endpoint to return Unicode Emoji data. This takes the Unicode emoji text file and the CLDR annotations and maps each emoji to its codepoint(s), character, name, and search terms.

An example of the data is:
```json
{
    "Smileys & Emotions": [
        {
            "character": "😀",
            "codepoints": "1F600",
            "annotations": [
                "face",
                "grin",
                "grinning face"
            ],
            "name": "grinning face"
        },
        {
            "character": "😃",
            "codepoints": "1F603",
            "annotations": [
                "face",
                "grinning face with big eyes",
                "mouth",
                "open",
                "smile"
            ],
            "name": "grinning face with big eyes"
        },
        {
            "character": "😄",
            "codepoints": "1F604",
            "annotations": [
                "eye",
                "face",
                "grinning face with smiling eyes",
                "mouth",
                "open",
                "smile"
            ],
            "name": "grinning face with smiling eyes"
        },
        ...
        {
            "character": "👋",
            "codepoints": "1F44B",
            "annotations": [
                "hand",
                "wave",
                "waving"
            ],
            "name": "waving hand"
        },
        {
            "character": "👋🏻",
            "codepoints": "1F44B 1F3FB",
            "annotations": null,
            "name": "waving hand: light skin tone"
        },
        {
            "character": "👋🏼",
            "codepoints": "1F44B 1F3FC",
            "annotations": null,
            "name": "waving hand: medium-light skin tone"
        },
        {
            "character": "👋🏽",
            "codepoints": "1F44B 1F3FD",
            "annotations": null,
            "name": "waving hand: medium skin tone"
        },
        {
            "character": "👋🏾",
            "codepoints": "1F44B 1F3FE",
            "annotations": null,
            "name": "waving hand: medium-dark skin tone"
        },
        {
            "character": "👋🏿",
            "codepoints": "1F44B 1F3FF",
            "annotations": null,
            "name": "waving hand: dark skin tone"
        },
    ]
}
```