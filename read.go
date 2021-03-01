package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Declare struct and create an empty slice
	type Person struct {
		fname string
		lname string
	}

	var fileContent []Person

	// Prompt for a filename
	var filename, line string

	fmt.Printf("Please enter the name of the text file: ")
	fmt.Scan(&filename)

	// Open file and read lines
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line = scanner.Text()
		lineSlice := strings.Split(line, " ")
		fileContent = append(fileContent, Person{lineSlice[0], lineSlice[1]})
	}

	// Print result
	for _, person := range fileContent {
		fmt.Printf("- %s %s\n", person.fname, person.lname)
	}

}
