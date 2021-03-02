package main

import "fmt"

func main() {

	f1()
	f2()
}

var x int = 100

func f1() {
	x := 4
	fmt.Printf("f1() x = %d\n", x)
	fmt.Printf("f1() &x = %v\n\n", &x)
}

func f2() {
	fmt.Printf("f2() x = %d\n", x)
	fmt.Printf("f2() &x = %v\n", &x)
}
