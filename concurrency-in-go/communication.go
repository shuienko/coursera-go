package main

import (
	"fmt"
	"sync"
)

func sumElements(s []int, c chan int, group *sync.WaitGroup) {
	sum := 0
	for _, value := range s {
		sum += value
	}
	c <- sum
	group.Done()
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int)
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	wg.Add(2)
	go sumElements(array[:5], c, &wg)
	go sumElements(array[5:], c, &wg)

	go func() {
		for item := range c {
			fmt.Println(item)
		}
	}()

	wg.Wait()
}
