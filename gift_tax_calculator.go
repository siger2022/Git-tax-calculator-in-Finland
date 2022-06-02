//https://www.vero.fi/en/individuals/property/gifts/gift-tax-calculator/

package main

import (
	"fmt"
)

func main() {
	var gift_value, tax float32

	fmt.Print("Input the gift value: ")
	fmt.Scanln(&gift_value)

	if gift_value < 5000 {
		fmt.Println("No tax!")

	} else if gift_value >= 5_000 && gift_value < 25_000 {
		tax = 100 + (gift_value-5000)*0.08
		fmt.Printf("The amount of gift tax is %v", tax)

	} else if gift_value >= 25_000 && gift_value < 55_000 {
		tax = 1_700 + (gift_value-25_000)*0.1
		fmt.Printf("The amount of gift tax is %v", tax)

	} else if gift_value >= 55_000 && gift_value < 200_000 {
		tax = 4_700 + (gift_value-55_000)*0.12
		fmt.Printf("The amount of gift tax is %v", tax)

	} else if gift_value >= 200_000 && gift_value < 1_000_000 {
		tax = 22_100 + (gift_value-200_000)*0.15
		fmt.Printf("The amount of gift tax is %v", tax)

	} else if gift_value > 1_000_000 {
		tax = 142_100 + (gift_value-1_000_000)*0.17
		fmt.Printf("The amount of gift tax is %v", tax)
	}

}
