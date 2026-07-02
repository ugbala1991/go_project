package main

import "fmt"

// Constants — values that never change
const (
	ShopName        = "BethyBuy" // const string
	TaxRate         = 0.075      //const float64
	MaxDiscount int = 50         //const int
)

func main() {

	// Variables — customer info
	customerName := "uche" //:= string
	itemsBought := 3       //:= int
	pricePerItem := 29.99  //:= float64
	isMember := false      // := bool (can as well change to true)

	// Calculate subtotal
	subtotal := float64(itemsBought) * pricePerItem

	// Apply discount if customer is a member
	var discount float64 = 0 //var float64
	if isMember {
		discount = 10.0 // flat ₦10 off for members
	}

	// Calculate tax and total
	taxAmount := (subtotal - discount) * TaxRate
	total := subtotal - discount + taxAmount

	// Print the receipt
	fmt.Println("=====", ShopName, "Receipt =====")
	fmt.Println("Customer:", customerName)
	fmt.Printf("Items: %d x $%.2f\n", itemsBought, pricePerItem)
	fmt.Printf("Subtotal: $%.2f\n", subtotal)
	fmt.Printf("Discount: -$%.2f\n", discount)
	fmt.Printf("Tax (%.1f%%): $%.2f\n", TaxRate*100, taxAmount)
	fmt.Printf("Total: $%.2f\n", total)
	fmt.Println("Member discount applied:", isMember)
}
