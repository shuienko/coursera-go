package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x [5]string
	fmt.Println(x, reflect.TypeOf(x))

	x[2] = "hello"
	fmt.Println(x, reflect.TypeOf(x))

	y := [3]int{1, 2, 3}
	fmt.Println(y, reflect.TypeOf(y))

	z := [...]int{11, 22, 33, 44}
	fmt.Println(z, reflect.TypeOf(z))

	for i, v := range z {
		fmt.Printf("index = %d | value = %d\n", i, v)
	}

	z1 := z[0:3]
	z2 := z[2:4]

	fmt.Println(z1, z2)
	fmt.Println(len(z1), cap(z1))
	fmt.Println(len(z2), cap(z2))

	z2[0] = 100
	fmt.Println(z1, z2)

	sli := []int{2, 3, 4, 5, 6}
	fmt.Println(sli, reflect.TypeOf(sli))

	m1 := make([]string, 5)
	fmt.Println(m1, reflect.TypeOf(m1))

	sli = append(sli, 7, 8, 9)
	fmt.Println(sli, reflect.TypeOf(sli))
}
