package main

import (
	"fmt"
	_ "packageA"
	_ "packageB"
	"time"
)

func main() {
	fmt.Println("main")
	time.Sleep(time.Second * 2)
}
