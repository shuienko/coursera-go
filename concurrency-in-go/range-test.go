package main

import (
	"fmt"
	"time"
)

func producer(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(time.Millisecond * 300)
	}
	close(c)
}

func main() {
	c := make(chan int)
	go producer(c)

	for item := range c {
		fmt.Println(item)
	}
}
