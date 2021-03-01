package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	nameAddr := make(map[string]string)

	// Get a name from uset
	fmt.Printf("Please, enter a name: ")
	in := bufio.NewReader(os.Stdin)
	input1, _ := in.ReadString('\n')

	// Get an address from user
	fmt.Printf("Please, enter an address: ")
	input2, _ := in.ReadString('\n')

	// Save everything to the map
	nameAddr["name"] = strings.TrimSuffix(input1, "\n")
	nameAddr["address"] = strings.TrimSuffix(input2, "\n")

	// Create JSON and print it
	jsonOutput, err := json.Marshal(nameAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("JSON object:", string(jsonOutput))
}
