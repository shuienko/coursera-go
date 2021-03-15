package main

import (
	"fmt"
	"time"
)

var swtch bool

var n int

func incN() {
	for i := 0; i < 10000; i++ {
		n++
	}
	swtch = true
}

func printN() {
	fmt.Println("printN() ->", n)
}

func printNSwitch() {
	for {
		if swtch == true {
			fmt.Println("printNSwitch() ->", n)
			break
		}
	}
}

func main() {
	go incN()
	go printNSwitch()
	go printN()

	time.Sleep(time.Second)
}
