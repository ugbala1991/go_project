package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("BANNER options: standard, shadow, thinkertoy (default: standard)")
		return
	}
	//handle os config
	userInput := os.Args[1]
	bannerFile := "standard.txt"
	if len(os.Args) == 3 {
		bannerFile = os.Args[2]
	}

	//handle empty string
	if userInput == "" {
		return
	}

	// 	// Handle the special case of just a newline
	if userInput == "\\n" {
		fmt.Println()
		return
	}

	// Replace literal \n strings with actual newline characters
	userInput = strings.ReplaceAll(userInput, "\\n", "\n")

	// Load the banner (defaulting to standard.txt)
	bannerData, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	//check when banner is empty
	if len(bannerData) == 0 {
		fmt.Println("empty banner file")
		return
	}

	PrintASCII(userInput, string(bannerData))
}
