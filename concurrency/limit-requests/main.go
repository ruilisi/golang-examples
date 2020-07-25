package main

import (
	"fmt"
	"time"
)

const MAXREQS = 50

var sem = make(chan int, MAXREQS)

type Request struct {
	a, b   int
	replyc chan int
}

var requestCount = 0

func process() {
	time.Sleep(time.Second)
	requestCount++
	fmt.Printf("Request %d finished\n", requestCount)
}

func handle() {
	sem <- 1 // doesn't matter what we put in it
	process()
	<-sem // one empty place in the buffer: the next request can start
}

func server(service chan *Request) {
	for {
		go handle()
	}
}

func main() {
	service := make(chan *Request)
	go server(service)
	time.Sleep(time.Minute)
}
