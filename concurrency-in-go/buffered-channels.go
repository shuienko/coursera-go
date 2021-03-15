package main

import (
	"fmt"
	"time"
)

func generator(c chan int) {
	n := 0
	for {
		c <- n
		n++
	}
}

func receiver(c chan int) {
	for {
		for i := 0; i < len(c); i++ {
			fmt.Println(<-c)
		}
		fmt.Println("-----")
		time.Sleep(time.Second)
	}
}

func main() {
	c := make(chan int, 5)

	go generator(c)
	go receiver(c)

	time.Sleep(time.Second * 20)
}
