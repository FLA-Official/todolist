package utils

import "strings"

func GeneratePrefix(title string) string {
	words := strings.Fields(title)

	var prefix string
	for _, w := range words {
		prefix += strings.ToUpper(string(w[0]))
	}

	return prefix
}
