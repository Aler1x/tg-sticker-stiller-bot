package utils

import (
	"regexp"
	"strings"
)

var (
	stickerPackRegex = regexp.MustCompile("(?:https?://)?(?:www.)?t.me/addstickers/([a-zA-Z0-9_]+)")
	emojiPackRegex   = regexp.MustCompile("(?:https?://)?(?:www.)?t.me/addemoji/([a-zA-Z0-9_]+)")
)

func IsStickerPack(text string) bool {
	return stickerPackRegex.MatchString(text)
}

func IsEmojiPack(text string) bool {
	return emojiPackRegex.MatchString(text)
}

func ExtractStickerPackName(text string) string {
	matches := stickerPackRegex.FindStringSubmatch(text)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func ExtractEmojiPackName(text string) string {
	matches := emojiPackRegex.FindStringSubmatch(text)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func NormalizePackName(input string) string {
	// Convert to lowercase
	normalized := regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(
		regexp.MustCompile(`\s+`).ReplaceAllString(strings.ToLower(strings.TrimSpace(input)), "_"),
		"_",
	)

	// Remove leading/trailing underscores
	normalized = strings.Trim(normalized, "_")

	// Replace multiple consecutive underscores with single underscore
	normalized = regexp.MustCompile(`_+`).ReplaceAllString(normalized, "_")

	return normalized
}

func ValidateNormalizedName(name string) bool {
	if len(name) == 0 || len(name) > 64 {
		return false
	}

	// Only lowercase letters, digits, and underscores
	validNameRegex := regexp.MustCompile(`^[a-z0-9_]+$`)
	return validNameRegex.MatchString(name)
}

func GetValidationError(name string) string {
	if len(name) == 0 {
		return "name-empty"
	}
	if len(name) > 64 {
		return "name-too-long"
	}

	return ""
}
