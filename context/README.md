## context

一种控制并发的方式

### [with-cancel.go](with-cancel.go)

一个监控并发，当 `context cancel()` 时，退出这个进程，结束

### [backgroud.go](backgroud.go) 

三个进程同时监控，十秒之后 `context cancel()` ，三个监控并发关闭

### 总结

`context cancel()` 可以与并发相结合，控制并发的结束，实现同步






