package main

import "fmt"

func getClosure() func() {
	count := 0
	return func() {
		count++
		fmt.Println(count)
	}
}

func main() {
	counter := getClosure()
	counter()
	counter()
	counter()
	counter()
	counter()
	counter()
	counter()
	counter()
}
