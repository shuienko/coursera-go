package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var wgr sync.WaitGroup

func printer() {
	once.Do(func() {
		fmt.Println("ololo")
	})
	fmt.Println("I'm printer")
	wgr.Done()
}

func main() {
	wgr.Add(3)
	go printer()
	go printer()
	go printer()
	wgr.Wait()
}
