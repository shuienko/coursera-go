/* EXPLANATION!

triangleArea() computes area based on global variables 'a' and 'h'
setTriangleParameters() sets 'a' and 'h' one by one.

Without concurrency everything is ok:
1. [setTriangleParameters()] set 'a' to some non-zero value
2. [setTriangleParameters()] set 'h' to some non-zero value
3. [triangleArea()] calculate area 'S'
4. [triangleArea()] print out result

But with concurrency enables we can potentially get:
1. [setTriangleParameters()] set 'a' to some non-zero value
2. [triangleArea()] calculate area 'S'
...

^^^ Here is the issue. Second parameter 'h' will remain with it's zero value - '0'.
As a result Triangle Area is not correct and equals to 0. Just because order of execution.
*/

package main

import (
	"fmt"
	"time"
)

// Global variables for triangle params
var a, h int

// Calculate Triangle Area based on global variables 'a' and 'h'
func triangleArea() {
	S := a * h / 2
	fmt.Println("S =", S)
}

// Set Global variables 'a' and 'h'
func setTriangleParameters(base, height int) {
	a = base
	h = height
}

func main() {
	go setTriangleParameters(8, 6)
	triangleArea()

	fmt.Println("Done!")
	time.Sleep(time.Second)
}
