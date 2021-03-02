package main

import (
	"fmt"
	"reflect"
)

func main() {
	m1 := map[int]string{
		10: "Alex",
		12: "Jane",
		33: "Bob",
	}

	fmt.Println(m1[10], reflect.TypeOf(m1))

	var m2 map[int]int
	m2 = make(map[int]int)
	fmt.Println(m2, m2[3], reflect.TypeOf(m2))

	for k, v := range m1 {
		fmt.Printf("key = %d | value = %s\n", k, v)
	}

	delete(m1, 33)
	fmt.Println(m1)

	name, exists := m1[12]
	fmt.Println(name, exists)

}
