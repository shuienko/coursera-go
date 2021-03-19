package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func randGen(c chan map[string]int, tag string, d int, g *sync.WaitGroup) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	i := 0
	for {
		c <- map[string]int{tag: i}
		i++

		delay := r.Intn(d)
		time.Sleep(time.Millisecond * time.Duration(delay))
	}
	g.Done()
}

func selectReceive(c1, c2 chan map[string]int, g *sync.WaitGroup) {
	for {
		select {
		case v1 := <-c1:
			fmt.Println(v1)
		case v2 := <-c2:
			fmt.Println(v2)
		}
	}

	g.Done()
}

func main() {
	var wg sync.WaitGroup
	c1 := make(chan map[string]int)
	c2 := make(chan map[string]int)

	wg.Add(2)
	go randGen(c1, "one", 1000, &wg)
	go randGen(c2, "two", 3000, &wg)
	go selectReceive(c1, c2, &wg)
	wg.Wait()

}
