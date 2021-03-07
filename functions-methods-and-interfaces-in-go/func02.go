package main

import "fmt"

type Fun func() string

func returnFunc(n int) func() string {
	if n < 0 {
		return func() string {
			return "n is less than 0"
		}
	} else {
		return func() string {
			return "n is equals of greater than 0"
		}
	}
}

func Average(numbers ...int) float64 {
	var sum int
	for _, n := range numbers {
		sum += n
	}

	return float64(sum) / float64(len(numbers))
}

func main() {
	var f1 Fun = func() string {
		return "I am f1"
	}

	var f2 func() = func() {
		fmt.Println("I am f2")
	}

	fmt.Println(f1())
	f2()

	fmt.Println(returnFunc(5)())

	numbers := []int{2, 5, 7, 9}
	fmt.Printf("%.6f\n", Average(1, 2, 4, 6))
	fmt.Printf("%.6f\n", Average(numbers...))
}
