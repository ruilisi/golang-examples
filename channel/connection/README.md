## connection

**[connetion1.go](connection1.go)**

**无缓存的Channel上的发送操作总在对应的接收操作完成前发生.**

```go
package main

var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "你好, 世界"
	done <- true
}

func main() {
	go aGoroutine()
	<-done
	println(msg)
}

```

**[connetion2.go](connection2.go)**

该程序首先对`msg`进行写入，然后在`done`管道上发送同步信号，随后从`done`接收对应的同步信号，最后执行`println`函数。

若在关闭Channel后继续从中接收数据，接收者就会收到该Channel返回的零值。

```go
package main

var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "你好, 世界"
	close(done)
}

func main() {
	go aGoroutine()
	<-done
	println(msg)
}

```

