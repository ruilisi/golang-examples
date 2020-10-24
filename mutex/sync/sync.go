package main

import "time"

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
	time.Sleep(time.Second * 10)
}

func main() {
	go setup()
	for !done {
	}
	print(a)
}
