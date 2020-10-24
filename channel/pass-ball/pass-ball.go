package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")

	myBall := 1314
	ch := make(chan int)
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		ball := <-ch
		fmt.Println("world")
		ch1 <- ball
	}()

	go func() {
		ball := <-ch2
		fmt.Println("you")
		ch3 <- ball
	}()

	go func() {
		ball := <-ch1
		fmt.Println("I love")
		ch2 <- ball
	}()

	ch <- myBall
	fmt.Println(<-ch3)
}
