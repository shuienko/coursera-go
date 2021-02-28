package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var x string = "he1lo"
	var y string = "22"
	fmt.Println(x[2], reflect.TypeOf(x[2]))

	fmt.Println("is digit? ", unicode.IsDigit(rune(x[2])))
	fmt.Println("is letter? ", unicode.IsLetter(rune(x[2])))

	fmt.Println(strings.Count(x, "l"))
	fmt.Println(strings.Replace(x, "l", "L", -1))

	var yInt, err = strconv.Atoi(y)
	if err != nil {
		fmt.Println("Can't convert: ", err)
	}
	fmt.Println(yInt, reflect.TypeOf(yInt))

}
