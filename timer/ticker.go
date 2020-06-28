package main

import (
	"fmt"
	"time"
)

func main() {
	//初始化断续器,间隔2s
	var ticker *time.Ticker = time.NewTicker(1 * time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Second * 5) //阻塞，则执行次数为sleep的休眠时间/ticker的时间
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
