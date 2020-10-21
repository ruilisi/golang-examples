## 超时机制


线程休眠5秒

```
	go func() {
		time.Sleep(time.Second * 5)
		ch <- 100
	}()
```

给10秒中的时间，如果10秒钟之内有响应，输出响应的结果，否则输出超时，结束


```
	select {
	case resp := <-ch:
		fmt.Println(resp)
	case <-time.After(time.Second * 10):
		fmt.Println("timeout")
		break
	}
```
