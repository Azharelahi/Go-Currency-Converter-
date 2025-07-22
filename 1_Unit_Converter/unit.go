package main

import (
	"fmt"
)

func main() {

	const (
		PKR_TO_USD = 1 / 278.5
		USD_TO_PKR = 278.5
		USD_TO_EUR = 0.91
		EUR_TO_USD = 1 / 0.91
		PKR_TO_EUR = PKR_TO_USD * USD_TO_EUR
		EUR_TO_PKR = 1 / PKR_TO_EUR
	)

	currencies := []string{"PKR", "USD", "EUR"}

	fmt.Println("💱 Available Currencies:")
	for i, currency := range currencies {
		fmt.Printf("%d. %s\n", i+1, currency)
	}

	fmt.Print("\nChoose your source currency (1–3): ")
	var source int
	fmt.Scanln(&source)

	if source < 1 || source > 3 {
		fmt.Println("Invalid source currency selected.")
		return
	}

	fmt.Println("\nSelect the currency you want to convert to:")
	for i := 0; i < len(currencies); i++ {
		if i+1 != source {
			fmt.Printf("%d. %s\n", i+1, currencies[i])
		}
	}
	fmt.Print("Enter your target currency (1–3): ")
	var target int
	fmt.Scanln(&target)

	if target < 1 || target > 3 || target == source {
		fmt.Println("Invalid target currency selected.")
		return
	}

	var amount float64
	fmt.Print("\nEnter the amount to convert: ")
	fmt.Scanln(&amount)

	var result float64

	
	switch currencies[source-1] + "_" + currencies[target-1] {
	case "PKR_USD":
		result = amount * PKR_TO_USD
	case "USD_PKR":
		result = amount * USD_TO_PKR
	case "USD_EUR":
		result = amount * USD_TO_EUR
	case "EUR_USD":
		result = amount * EUR_TO_USD
	case "PKR_EUR":
		result = amount * PKR_TO_EUR
	case "EUR_PKR":
		result = amount * EUR_TO_PKR
	default:
		fmt.Println("Conversion not supported.")
		return
	}

	fmt.Printf("\n%.2f %s = %.2f %s\n",
		amount, currencies[source-1],
		result, currencies[target-1])
}
