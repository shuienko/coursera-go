package main

import (
	"fmt"
	"math"
)

func main() {
	var floatNumber float64

	// Get a floating point number from user
	fmt.Printf("Enter a floating point number: ")
	_, err := fmt.Scan(&floatNumber)
	if err != nil {
		fmt.Println(err)
	}

	// Truncate to integer and print result
	intNumber := int(math.Trunc(floatNumber))
	fmt.Println(intNumber)
}
