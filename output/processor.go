package main

import (
	"strings"
)

// PrintASCII generates the ASCII art and returns it as a string.
func PrintASCII(input, banner string) string {
	bannerLines := strings.Split(strings.ReplaceAll(banner, "\r\n", "\n"), "\n")

	if len(bannerLines) < 855 {
		return "incomplete banner\n"
	}

	var result strings.Builder

	words := strings.Split(input, "\n")

	for _, word := range words {
		if word == "" {
			result.WriteString("\n")
			continue
		}

		for i := 1; i <= 8; i++ {
			for _, char := range word {
				if char < 32 || char > 126 {
					continue
				}

				start := int(char-32)*9 + i

				if start >= 0 && start < len(bannerLines) {
					result.WriteString(bannerLines[start])
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String()
}