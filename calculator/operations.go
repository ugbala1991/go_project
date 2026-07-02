package main

import (
	"fmt"
	"strconv"
)

// OperationFunc defines the type for calculator operations
type OperationFunc func(args []string) error

// Helper to parse two numbers safely
func parseTwoNumbers(args []string) (float64, float64, error) {
	if len(args) != 2 {
		return 0, 0, fmt.Errorf("expected 2 arguments, got %d", len(args))
	}
	a, err1 := strconv.ParseFloat(args[0], 64)
	b, err2 := strconv.ParseFloat(args[1], 64)
	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("arguments must be valid numbers")
	}
	return a, b, nil
}

func add(args []string) error {
	a, b, err := parseTwoNumbers(args)
	if err != nil {
		return err
	}
	fmt.Printf("✦ Result: %.2f\n", a+b)
	return nil
}

func sub(args []string) error {
	a, b, err := parseTwoNumbers(args)
	if err != nil {
		return err
	}
	fmt.Printf("✦ Result: %.2f\n", a-b)
	return nil
}

func mul(args []string) error {
	a, b, err := parseTwoNumbers(args)
	if err != nil {
		return err
	}
	fmt.Printf("✦ Result: %.2f\n", a*b)
	return nil
}

func div(args []string) error {
	a, b, err := parseTwoNumbers(args)
	if err != nil {
		return err
	}
	if b == 0 {
		return fmt.Errorf("division by zero is not allowed")
	}
	fmt.Printf("✦ Result: %.2f\n", a/b)
	return nil
}
