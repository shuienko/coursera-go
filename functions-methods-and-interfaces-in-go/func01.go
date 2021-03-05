package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sayHello() {
	fmt.Println("Hello!")
}

func TriangleArea(a, h float64) float64 {
	return a * h / 2
}

func RollTheDice() (int, int) {
	rand.Seed(time.Now().UnixNano())
	return (rand.Int() % 6) + 1, (rand.Int() % 6) + 1
}

func Increment(n *int) {
	*n++
}

func Sum(a []int) int {
	s := 0
	for _, item := range a {
		s += item
	}

	a[0] = 100

	return s
}

func main() {
	sayHello()
	fmt.Println("S =", TriangleArea(5, 8))
	one, two := RollTheDice()

	fmt.Println(one, two)

	x := 1
	Increment(&x)
	fmt.Println(x)

	Increment(&x)
	fmt.Println(x)

	Increment(&x)
	fmt.Println(x)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(Sum(arr), arr)

}
