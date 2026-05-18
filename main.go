package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("BANNER options: standard, shadow, thinkertoy (defualt : standard)")
	}

	userInput := os.Args[1]
	bannerType := "standard.txt"
	if len(os.Args) == 3 {
		userInput = os.Args[2] + ".txt"
	}

	if userInput == "" {
		return
	}

	if userInput == "\\n" {
		fmt.Println()
		return
	}

	userInput = strings.ReplaceAll(userInput, "\\n", "\n")

	bannerData, err := os.ReadFile(bannerType)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	PrintASCII(userInput, string(bannerData))
}
