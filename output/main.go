package main

import (
	"fmt"
	"os"
	"strings"
)

func usage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("EX: go run . --output=<fileName.txt> something standard")
}

func main() {
	var outputFile string

	args := os.Args[1:]

	for _, arg := range args {
	if strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--output=") {
		usage()
		return
	}
}

	// Check for valid number of arguments
	if len(args) == 0 || len(args) > 3 {
		usage()
		return
	}

	// Handle --output flag
	if strings.HasPrefix(args[0], "--output=") {
		outputFile = strings.TrimPrefix(args[0], "--output=")

		// Invalid format: --output=
		if outputFile == "" {
			usage()
			return
		}

		args = args[1:]
	}

	// After removing the flag, there should be 1 or 2 arguments left
	if len(args) == 0 || len(args) > 2 {
		usage()
		return
	}

	userInput := args[0]

	// Default banner
	bannerFile := "standard.txt"

if len(args) == 2 {
	switch args[1] {
	case "standard", "shadow", "thinkertoy":
		bannerFile = args[1] + ".txt"
	default:
		fmt.Println("invalid banner")
		return
	}
}
	// Handle empty input
	if userInput == "" {
		return
	}

	// Handle only "\n"
	if userInput == "\\n" {
		if outputFile != "" {
			os.WriteFile(outputFile, []byte("\n"), 0644)
		} else {
			fmt.Println()
		}
		return
	}

	// Replace literal "\n" with actual newlines
	userInput = strings.ReplaceAll(userInput, "\\n", "\n")

	// Read banner file
	bannerData, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Check for empty banner file
	if len(bannerData) == 0 {
		fmt.Println("empty banner file")
		return
	}

	// Generate ASCII art
	result := PrintASCII(userInput, string(bannerData))

	// Write to file if requested
	if outputFile != "" {
		err = os.WriteFile(outputFile, []byte(result), 0644)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		return
	}

	// Otherwise print to terminal
	fmt.Print(result)
}