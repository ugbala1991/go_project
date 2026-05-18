package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// -------Process Function------------
	// result := string(data)
	// result = UpLowCap(result)
	// result = BinConversion(result)
	// result = HexConversion(result)
	// result = Punctuation(result)
	// result = Quotuation(result)
	// result = fixAorAN(result)

	var resultLines []string
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		processed := BinConversion(line)
		processed = HexConversion(processed)
		processed = Punctuation(processed)
		processed = Quotuation(processed)
		processed = UpLowCap(processed)
		processed = fixAorAN(processed)

		resultLines = append(resultLines, processed)
	}

	result := strings.Join(resultLines, "\n")

	// Write output file
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing output:", err)
		return
	}

	fmt.Println("Processing complete. Check output.txt")
}
