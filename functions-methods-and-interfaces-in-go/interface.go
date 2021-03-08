package main

import (
	"fmt"
	"reflect"
)

type ITGuy interface {
	Complain()
	DrinkCoffee()
	SayHello()
}

type DevOps struct {
	name string
	age  int
}

type Programmer struct {
	name string
	age  int
}

func (d DevOps) Complain() {
	fmt.Println("Our old infra is so lame. Let's switch to Kubernetes!")
}

func (d DevOps) DrinkCoffee() {
	fmt.Println("I love this job!")
}

func (d DevOps) SayHello() {
	fmt.Printf("Hello! My name is %s\n", d.name)
}

func (p Programmer) Complain() {
	fmt.Println("It works on my machine!")
}

func (p Programmer) DrinkCoffee() {
	fmt.Println("Life is good!")
}

func (p Programmer) SayHello() {
	fmt.Printf("Hello! My name is %s\n", p.name)
}

func dailyRoutine(person ITGuy) {
	person.Complain()
	person.DrinkCoffee()
}

func getType(v interface{}) {
	fmt.Println(reflect.TypeOf(v))
}

func main() {
	d1 := DevOps{"Oleksandr", 32}
	p1 := Programmer{"Anton", 31}

	dailyRoutine(d1)
	fmt.Println("")
	dailyRoutine(p1)

	// nil dynamic value test
	var it1 ITGuy
	var d2 DevOps
	it1 = d2

	it1.SayHello() // returns zero value for name field.

	// nil dynamic type test
	//var it2 ITGuy
	//it2.SayHello() // exception

	// Type assertion examples
	getType(it1)
	d3, ok := it1.(DevOps)
	if ok {
		fmt.Println(d3)
		getType(d3)
	}

	switch p := it1.(type) {
	case DevOps:
		fmt.Println("DevOps", p)
	case Programmer:
		fmt.Println("Programmer", p)
	}
}
