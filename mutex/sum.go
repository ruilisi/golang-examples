package main

import (
	"fmt"
	"sync"
	"time"
)

var total struct {
	sync.Mutex
	value int
}

func worker(tName string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		total.Lock()
		time.Sleep(10 * time.Millisecond)
		fmt.Println(tName, i)
		total.value += i
		total.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker("alice", &wg)
	go worker("bob", &wg)
	wg.Wait()

	fmt.Println(total.value)
}
