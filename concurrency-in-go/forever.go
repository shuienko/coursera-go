package main

import (
	"fmt"
	"sync"
	"time"
)

func gen(c chan int, group *sync.WaitGroup) {
	i := 0
	for {
		c <- i
		i++
		time.Sleep(time.Second)
	}

}

func rcv(c chan int, group *sync.WaitGroup) {
	for {
		fmt.Println(<-c)
	}
	group.Done()
}

func main() {
	c := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go gen(c, &wg)
	go rcv(c, &wg)
	wg.Wait()
}
