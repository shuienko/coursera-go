package main

import (
	"fmt"
	"sync"
)

var N int = 10
var wg sync.WaitGroup

var mut sync.Mutex

func inc(tag string) {
	fmt.Println(tag)

	mut.Lock()
	for i := 0; i < 10000; i++ {
		N++
	}
	mut.Unlock()

	wg.Done()
}

func dec(tag string) {
	fmt.Println(tag)

	mut.Lock()
	for i := 0; i < 10000; i++ {
		N--
	}
	mut.Unlock()

	wg.Done()
}

func main() {
	fmt.Println("N =", N)
	wg.Add(5)
	go dec("dec-1")
	go dec("dec-2")
	go inc("inc-1")
	go inc("inc-2")
	go inc("inc-3")
	wg.Wait()
	fmt.Println("N =", N)
}
