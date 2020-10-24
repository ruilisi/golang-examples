package main

import (
	"fmt"
	"time"
)

func main() {
	//初始化断续器,间隔2s
	var ticker *time.Ticker = time.NewTicker(100 * time.Millisecond)

	//num为指定的执行次数
	num := 10
	c := make(chan int, num)
	go func() {
		for t := range ticker.C {
			c <- 1
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 10000)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
