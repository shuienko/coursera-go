package main

import (
	"fmt"
	"sync"
)

var number int

func incNumber(group *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		number++
	}
	defer group.Done()
}

func printNumber(group *sync.WaitGroup) {
	fmt.Println("printNumber() ->", number)
	defer group.Done()
}

func printNWait(group *sync.WaitGroup) {
	fmt.Println("printNWait() ->", number)
	defer group.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go incNumber(&wg)
	wg.Wait()

	wg.Add(2)
	go printNumber(&wg)
	go printNWait(&wg)
	wg.Wait()
}
