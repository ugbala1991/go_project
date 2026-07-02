package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convert(number string, base string) error {
	switch base {
	case "hex":
		// Parse hex (base 16)
		val, err := strconv.ParseInt(number, 16, 64)
		if err != nil {
			return fmt.Errorf("invalid hex number: %s", number)
		}
		fmt.Printf("✦ Decimal: %d\n", val)
	case "bin":
		// Parse binary (base 2)
		if !isValidBinary(number) {
			return fmt.Errorf("invalid binary number: %s", number)
		}
		val, err := strconv.ParseInt(number, 2, 64)
		if err != nil {
			return fmt.Errorf("invalid binary number: %s", number)
		}
		fmt.Printf("✦ Decimal: %d\n", val)
	case "dec":
		// Parse decimal (base 10)
		val, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid decimal number: %s", number)
		}
		fmt.Printf("✦ Binary:  %s\n", strconv.FormatInt(val, 2))
		fmt.Printf("✦ Hex:     %s\n", strings.ToUpper(strconv.FormatInt(val, 16)))
	default:
		return fmt.Errorf("unsupported base: %s. Use hex, bin, or dec", base)
	}
	return nil
}

func isValidBinary(s string) bool {
	for _, ch := range s {
		if ch != '0' && ch != '1' {
			return false
		}
	}
	return true
}

func printHelp() {
	fmt.Println("Supported commands:")
	fmt.Println("  convert <number> <base>")
	fmt.Println("    base options: hex, bin, dec")
	fmt.Println("    examples:")
	fmt.Println("      convert 1E hex   → Decimal: 30")
	fmt.Println("      convert 10 bin   → Decimal: 2")
	fmt.Println("      convert 255 dec  → Binary: 11111111, Hex: FF")
	fmt.Println("  help                 → show this help message")
	fmt.Println("  quit                 → exit the program")
}
