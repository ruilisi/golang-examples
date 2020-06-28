package main

import (
	"fmt"
	"time"
)

func main() {
	//初始化定时器
	t := time.NewTimer(2 * time.Second)
	//当前时间
	now := time.Now()
	fmt.Printf("Now time : %v.\n", now)

	expire := <-t.C
	fmt.Printf("Expiration time: %v.\n", expire)
}
