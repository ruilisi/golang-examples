package main

import (
	"fmt"
)

func work() {
	panic("hello")
}

func safelyDo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Work failed with %s", err)
		}
	}()
	work()
}

func main() {
	safelyDo()
}
