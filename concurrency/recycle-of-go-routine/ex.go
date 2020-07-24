package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 0)
	go func() {
		time.Sleep(time.Second * 5)
		ch <- 100
	}()
	select {
	case resp := <-ch:
		fmt.Println(resp)
	case <-time.After(time.Second * 10):
		fmt.Println("timeout")
		break
	}
}
