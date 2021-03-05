package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Swap the contents of the slice in position i with the contents in position i+1
func Swap(slice []int, index int) {
	tmp := slice[index+1]
	slice[index+1] = slice[index]
	slice[index] = tmp
}

// Get user input and return slice of integers
func GetUserInput() []int {
	// Read input
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter integer numbers separated by whitespace: ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSuffix(userInput, "\n")

	// Split input string into slice of strings
	stringSlice := strings.Split(userInput, " ")

	// Convert slice of strings to slice of ints
	var intSlice []int
	for _, number := range stringSlice {
		n, _ := strconv.Atoi(number)
		intSlice = append(intSlice, n)
	}

	return intSlice
}

// Sort the slice of integers
func BubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1; j++ {
			if slice[j+1] < slice[j] {
				Swap(slice, j)
			}
		}
	}
}

func main() {
	input := GetUserInput()
	BubbleSort(input)
	fmt.Println("Sorted:", input)
}
