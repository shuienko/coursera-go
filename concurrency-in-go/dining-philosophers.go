package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Stick struct {
	sync.Mutex
}

type Philosopher struct {
	leftStick, rightStick *Stick
	number                int
}

func (p Philosopher) eat(c chan int) {
	for i := 0; i < 3; i++ {

		// Get "permission".
		// If buffered channel 'c' is full then execution will be blocked automatically.
		c <- 1

		p.leftStick.Lock()
		p.rightStick.Lock()

		fmt.Println("starting to eat", p.number)
		fmt.Println("finishing eating", p.number)

		p.leftStick.Unlock()
		p.rightStick.Unlock()
	}
	wg.Done()
}

// Reads from channel with capacity 2 (defined below).
// Will act as a "host" blocking execution after buffer is getting full
func host(c chan int) {
	for {
		<-c
	}
}

func main() {
	// Buffered channel for the host() function
	c := make(chan int, 2)

	// Create chopsticks
	Sticks := make([]*Stick, 5)
	for i := 0; i < 5; i++ {
		Sticks[i] = new(Stick)
	}

	// Create philosophers
	Philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		Philosophers[i] = &Philosopher{leftStick: Sticks[i], rightStick: Sticks[(i+1)%5], number: i + 1}
	}

	// Start host
	go host(c)

	// Start dining
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go Philosophers[i].eat(c)
	}
	wg.Wait()

}
