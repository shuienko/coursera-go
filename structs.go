package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	p1 := Person{
		Name: "Alex",
		Age:  32,
		Sex:  "male",
	}

	var p2 Person

	fmt.Println(p1, p1.Age, reflect.TypeOf(p1))
	fmt.Println(p2, reflect.TypeOf(p2))

	p3 := new(Person)
	fmt.Println(p3, reflect.TypeOf(p3))
}
