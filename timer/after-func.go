package main

import (
	"fmt"
	"time"
)

func main() {
	var t *time.Timer

	f := func() {
		fmt.Printf("Expiration time : %v.\n", time.Now())
		fmt.Printf("C`s len: %d\n", len(t.C))
	}

	t = time.AfterFunc(1*time.Second, f)
	//让当前Goroutine 睡眠2s，确保大于内容的完整
	//这样做原因是，time.AfterFunc的调用不会被阻塞。它会以一部的方式在到期事件来临执行我们自定义函数f。
	time.Sleep(2 * time.Second)
}
