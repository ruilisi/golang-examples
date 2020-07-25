package main

import (
	"flag"
	"fmt"
	"time"
)

var ngoroutine = flag.Int("n", 100000, "how many goroutines")

func f(left, right chan int) {
	v := 1 + <-right
	fmt.Printf("Passing %d from right channel to left channel\n", v)
	left <- v
}

func main() {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	time.Sleep(time.Second * 2)
	right <- 0      // bang!
	x := <-leftmost // wait for completion
	fmt.Println(x)  // 100000, ongeveer 1,5 s
}
