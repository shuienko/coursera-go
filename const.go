package main

import (
	"fmt"
	"reflect"
)

const pi = 3.14

type englishLevels int

const (
	A1 englishLevels = iota
	A2
	B1
	B2
	C1
	C2
)

func main() {
	fmt.Println(pi)
	fmt.Println(B2, reflect.TypeOf(B2))
}
