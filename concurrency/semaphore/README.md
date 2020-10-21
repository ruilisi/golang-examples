## semaphore 同步

给 `channel` 发一个信号存在里面

`s.Signal()`

```
// release n resources
func (s semaphore) V(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}
```

`s.wait()`

解锁信号

```
/* signal-wait */
func (s semaphore) Wait(n int) {
	s.P(n)
}
```

执行解锁

```
// acquire n resources
func (s semaphore) P(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}
```






