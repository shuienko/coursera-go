package main

import "fmt"

func main() {
	var x int = 100
	var y int = 200

	var z *int
	z = &x

	fmt.Println("x = ", x, "y = ", y, "z = ", z)
	fmt.Println("&x = ", &x, " &y = ", &y, "*z = ", *z)

	i := new(float64)
	var j *float64 = i

	*j = 314

	fmt.Println("i = ", i, "j = ", j, "*j = ", *j, "*i = ", *i)
}
