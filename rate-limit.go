package main

import (
	"fmt"
	"time"
)

func main() {
	rate := time.Second / 5
	throttle := time.Tick(rate)
	for i := 0; i < 100; i++ {
		<-throttle // rate limit our Service.Method RPCs
		fmt.Printf("%s: %d\n", time.Now().String(), i)
	}
}
