package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	go func() {
		fmt.Println("Async work doing")
		time.Sleep(time.Second * 3)
		fmt.Println("Async work done")
		close(done)
	}()
	<-done
	fmt.Println("Exit")
}
