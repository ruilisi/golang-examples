## [after-func.go](after-func.go)

让当前Goroutine 睡眠2s，确保大于内容的完整

 这样做原因是，time.AfterFunc的调用不会被阻塞。它会以一部的方式在到期事件来临执行我们自定义函数f。

`time.Sleep(2 * time.Second)`

## [after.go](after.go)

ch1，ch2中始终没有数据，因而最终一定会超时

```
ch1 := make(chan int, 1)
ch2 := make(chan int, 1)

select {
case e1 := <-ch1:
//如果ch1通道成功读取数据，则执行该case处理语句
fmt.Printf("1th case is selected. e1=%v", e1)
case e2 := <-ch2:
//如果ch2通道成功读取数据，则执行该case处理语句
fmt.Printf("2th case is selected. e2=%v", e2)
case <-time.After(2 * time.Second):
fmt.Println("Timed out")
}
```

## [reset.go](reset.go)

初始时的channel

```
ch11 := make(chan int, 1000)
sign := make(chan byte, 1)
```

如果ch11中有值，则输出ch11中的值

否者输出 `timeout` 结束

```
	for {
			select {
			case e = <-ch11:
				fmt.Printf("ch11 -> %d\n", e)
			case <-func() <-chan time.Time {
				if timer == nil {
					//初始化到期时间据此间隔1ms的定时器
					timer = time.NewTimer(time.Millisecond)
				} else {
					//复用，通过Reset方法重置定时器
					timer.Reset(time.Millisecond)
				}
				//得知定时器到期事件来临时，返回结果
				return timer.C
			}():
				fmt.Println("Timeout.")
				ok = false
				break
			}
			//终止for循环
			if !ok {
				sign <- 0
				break
			}
		}

	}()
```



## [simple.go](simple.go)

t是获得当前时间两秒钟之后的值

因而两个输出之间的值相差2秒

```
	t := time.NewTimer(2 * time.Second)
	//当前时间
	now := time.Now()
	fmt.Printf("Now time : %v.\n", now)

	expire := <-t.C
	fmt.Printf("Expiration time: %v.\n", expire)
```



## [ticker.go](ticker.go)

Ticker间隔为1秒，不断间断输出，总计输出5秒

```
	//初始化断续器,间隔1s
	var ticker *time.Ticker = time.NewTicker(1 * time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Second * 5) //阻塞，则执行次数为sleep的休眠时间/ticker的时间
	ticker.Stop()
	fmt.Println("Ticker stopped")
```





## [ticker-channel.go](ticker-channel.go)

num 大小为十，所以总共只能输出十个

`time.sleep()` 为线程睡眠时长

睡眠结束完成之后，结束

```
	//初始化断续器,间隔0.1s
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
```

