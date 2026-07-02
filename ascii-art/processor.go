package main

import (
	"fmt"
	"strings"
)

// PrintASCII splits the input by newlines and prints each segment line by line.
func PrintASCII(input, banner string) {
	// Standardize banner line endings and split by newline
	bannerLines := strings.Split(strings.ReplaceAll(banner, "\r\n", "\n"), "\n")

	//check incomplete banner file
	if len(bannerLines) < 855 {
		fmt.Println("incomplete banner")
		return
	}

	words := strings.Split(input, "\n")

	for _, word := range words {
		if word == "" {
			fmt.Println()
			continue
		}

		// Each character in the banner is 8 lines tall.
		// We loop 8 times to print the word row by row.
		for i := 1; i <= 8; i++ {
			line := ""

			for _, char := range word {
				// for _, char := range line {
				// THE LIFESAVER: Skip characters outside the banner range (ASCII 32-126)
				if char < 32 || char > 126 {
					continue
				}
				// Formula: (char - 32) * 9 + 1
				// ASCII ' ' is 32. In standard.txt, ' ' starts at line 1.
				// Each block is 9 lines (8 lines of art + 1 separator).
				start := int(char-32)*9 + i

				if start >= 0 && start < len(bannerLines) {
					line += bannerLines[start]
				}
			}
			fmt.Println(line)
		}
	}
}
