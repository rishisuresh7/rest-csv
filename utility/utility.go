package utility

import "strings"

func CheckList(list []string, key string) bool {
	for _, value := range list {
		if value == key {
			return true
		}
	}

	return false
}

func SanitizeFileName(file string) string {
	cleanName := strings.ReplaceAll(strings.TrimSpace(file), " ", "_")

	return strings.ToLower(cleanName)
}
