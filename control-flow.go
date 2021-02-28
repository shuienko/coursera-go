package main

import "fmt"

func main() {
	x := 5

	// IF
	if x >= 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than 5")
	}

	// FOR
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// SWITCH
	switch x {
	case 5:
		fmt.Println("x = 5")
	case 12:
		fmt.Println("x = 12")
	default:
		fmt.Println("x is something else")
	}

	// TAGLESS SWITCH
	switch {
	case x < 5:
		fmt.Println("x < 5")
	case x == 5:
		fmt.Println("x = 5")
	case x > 5:
		fmt.Println("x > 5")
	default:
		fmt.Println("WTF is x?")
	}

	var appleNum int
	fmt.Printf("Type a number, please: ")
	_, err := fmt.Scan(&appleNum)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(appleNum)
}
