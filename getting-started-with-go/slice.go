package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var slice []int
	var userInput string

	for {
		// Prompt user for input and save as a variable
		fmt.Printf("Please, enter an integer: ")
		_, err := fmt.Scan(&userInput)
		if err != nil {
			fmt.Println(err)
		}

		// Exit on "X" instead of integer
		if strings.ToUpper(userInput) == "X" {
			break
		}

		//Add integer to a slice, sort It
		number, _ := strconv.Atoi(userInput)
		slice = append(slice, number)
		sort.Ints(slice)

		// Print output
		fmt.Println(slice)
	}
}
