package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan string)
	go func() {
		for {
			requests <- "New request"
		}
	}()
	chRate := time.Tick(time.Second)
	for req := range requests {
		<-chRate
		fmt.Println(req)
	}
}
