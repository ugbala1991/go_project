package main

import (
	"fmt"
	"strings"
)

func PrintASCII(input, banner string) {
	bannerLines := strings.Split(strings.ReplaceAll(banner, "\r\n", "\n"), "\n")

	words := strings.Split(input, "\n")

	for _, word := range words {
		if word == "" {
			fmt.Println()
			continue
		}
		for i := 1; i <= 8; i++ {
			line := ""

			for _, char := range word {
				if char < 32 || char > 126 {
					continue
				}
				start := int(char-32)*9 + i
				if start >= 0 && start < len(bannerLines) {
					line += bannerLines[start]
				}

			}
			fmt.Println(line)
		}
	}

}
