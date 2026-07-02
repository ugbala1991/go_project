package main

import "fmt"

func help(args []string) error {
	fmt.Println("Supported commands:")
	fmt.Println("  add <a> <b>   → addition")
	fmt.Println("  sub <a> <b>   → subtraction")
	fmt.Println("  mul <a> <b>   → multiplication")
	fmt.Println("  div <a> <b>   → division")
	fmt.Println("  help          → show this help message")
	fmt.Println("  quit          → exit the calculator")
	return nil
}

func quit(args []string) error {
	fmt.Println("Goodbye! ✦")
	return nil
}
