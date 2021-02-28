package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Prompt user and read a string
	// bufio instead of fmt.Scan in order to handle strings with whitespaces
	fmt.Printf("Enter a string: ")
	in := bufio.NewReader(os.Stdin)
	userString, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	// Trim trailing newline character and switch all to lower case
	userStringLower := strings.ToLower(strings.TrimSuffix(userString, "\n"))

	// Check conditions and print result
	if strings.HasPrefix(userStringLower, "i") &&
		strings.HasSuffix(userStringLower, "n") &&
		strings.Contains(userStringLower, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
