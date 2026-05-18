package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// --------UpLowCap, 2-----------
func UpLowCap(text string) string {
	word := strings.Fields(text)
	for i := 0; i < len(word); i++ {
		if strings.HasPrefix(word[i], "(") && !strings.HasSuffix(word[i], ")") {
			mod := word[i] + word[i+1]
			mod = strings.Trim(mod, "()")
			modify := strings.Split(mod, ",")
			com := modify[0]
			n, _ := strconv.Atoi(modify[1])
			for j := i - n; j <= i; j++ {
				switch com {
				case "up":
					word[j] = strings.ToUpper(word[j])

				case "low":
					word[j] = strings.ToLower(word[j])

				case "cap":
					word[j] = strings.Title(strings.ToLower(word[j]))
				}
			}
			word = append(word[:i], word[i+2:]...)
			i--
		}

		if word[i] == "(up)" && i > 0 {
			word[i-1] = strings.ToUpper(word[i-1])
			word = append(word[:i], word[i+1:]...)
			i--
		}
		if word[i] == "(low)" && i > 0 {
			word[i-1] = strings.ToLower(word[i-1])
			word = append(word[:i], word[i+1:]...)
			i--
		}
		if word[i] == "(cap)" && i > 0 {
			word[i-1] = strings.Title(strings.ToLower(word[i-1]))
			word = append(word[:i], word[i+1:]...)
			i--
		}
	}
	return strings.Join(word, " ")
}

func HexConversion(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		word := words[i]

		//handle Hex
		if word == "(hex)" && i > 0 {
			val, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				words[i-1] = fmt.Sprintf("%d", val)
				words[i] = ""
			}
		}
	}
	return strings.Join(words, " ")
}

func BinConversion(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		word := words[i]

		//handle Hex
		if word == "(bin)" && i > 0 {
			val, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				words[i-1] = fmt.Sprintf("%d", val)
				words[i] = ""
			}
		}
	}
	return strings.Join(words, " ")
}

func fixAorAN(text string) string {
	// Case 1: "a" directly before a vowel/h
	re1 := regexp.MustCompile(`\b([Aa])\s+([aeiouhAEIOUH])`)
	text = re1.ReplaceAllString(text, "${1}n ${2}")

	// Case 2: "a" before a quote then vowel/h (e.g. a 'orange')
	re2 := regexp.MustCompile(`\b([Aa])\s+(['"])\s*([aeiouhAEIOUH])`)
	text = re2.ReplaceAllString(text, "${1}n ${2}${3}")

	return text
}

func Punctuation(text string) string {
	text = regexp.MustCompile(`\s+([,.?;:!])`).ReplaceAllString(text, `$1`)
	text = regexp.MustCompile(`([.,?;:!])(\s*)(\w)`).ReplaceAllString(text, "$1 $3")
	return text
}

func Quotuation(text string) string {
	//SingleQuote
	text = regexp.MustCompile(`'\s*(.*?)\s*'`).ReplaceAllString(text, ` '$1' `)
	//DoubleQuote
	text = regexp.MustCompile(`"\s*(.*?)\s*"`).ReplaceAllString(text, ` "$1" `)

	return text
}

// Fix "a" vs "an"
// func fixAorAN(text string) string {
//  words := strings.Fields(text)
//  for i := 0; i < len(words)-1; i++ {
//      if strings.ToLower(words[i]) == "a" {
//          next := strings.ToLower(words[i+1])
//          if strings.ContainsAny(string(next[0]), "aeiouh") {
//              words[i] = "an"
//          }
//      }
//  }
//  return strings.Join(words, " ")
// }
