package ascii

import (
	"errors"
	"os"
	"strings"
)

var ErrInvalidBanner = errors.New("invalid banner")

func GenerateASCII(text, banner string) (string, error) {
	if text == "" {
		return "", nil
	}

	switch banner {
	case "standard", "shadow", "thinkertoy":
	default:
		return "", ErrInvalidBanner
	}

	path := "banners/" + banner + ".txt"

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(data), "\n")
	textLines := strings.Split(text, "\n")

	var result strings.Builder

	for _, textLine := range textLines {

		if textLine == "" {
			result.WriteString("\n")
			continue
		}

		for row := 0; row < 8; row++ {

			for _, ch := range textLine {

				if ch < 32 || ch > 126 {
					continue
				}

				position := int(ch) - 32
				start := position * 9

				if start+row >= len(lines) {
					return "", errors.New("corrupted banner file")
				}

				result.WriteString(lines[start+row])
			}

			result.WriteString("\n")
		}
	}

	return result.String(), nil
}
