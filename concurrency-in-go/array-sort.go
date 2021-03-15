package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getUserInput() []int {
	// Read input
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter numbers separated by whitespace: ")
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

// Sort slice of integers and send result to the channel
func sortSlice(a []int, c chan []int) {
	fmt.Println("Processing:", a)

	sort.Ints(a)
	c <- a
}

// For a given slice of integers returns 4 sub-slices as a map in a format: map[begin]end
func partition4(a []int) map[int]int {
	var j int
	parts := make(map[int]int)
	chunkSize := len(a) / 4

	for i := 0; i < len(a); i += chunkSize {
		j += chunkSize
		if len(a)-i < chunkSize {
			continue
		}
		if len(a)-j < chunkSize {
			j = len(a)
		}
		parts[i] = j
	}
	return parts
}

func main() {
	c := make(chan []int, 4)
	arraySorted := []int{}

	// Get data from user and split array
	array := getUserInput()
	subSlices := partition4(array)

	// Start 4 goroutines
	for begin, end := range subSlices {
		go sortSlice(array[begin:end], c)
	}

	// Read from channel. Add results to the sorted array
	for i := 1; i <= 4; i++ {
		arraySorted = append(arraySorted, <-c...)
	}

	// Sort "sorted" chunks in array
	sort.Ints(arraySorted)

	// Print result
	fmt.Println("Sorted:", arraySorted)
}
