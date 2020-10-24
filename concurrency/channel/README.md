## 缓冲的channel

对于带缓冲的Channel，对于Channel的第K个接收完成操作发生在第K+C个发送操作完成之前，其中C是Channel的缓存大小。虽然管道是带缓存的，`main`线程接收完成是在后台线程发送开始但还未完成的时刻，此时打印工作也是已经完成的。

#### 使用channel等待线程完成 [channel1.go](channel2.go)

```go
package main

import "fmt"

func main() {
	done := make(chan int, 10) // 带 10 个缓存

	// 开N个后台打印线程
	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Println("你好, 世界")
			done <- 1
		}()
	}

	// 等待N个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}
}

```

#### 借助 `sync.WaitGroup`等待线程完成 [channel2.go](channel2.go)

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// 开N个后台打印线程
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			fmt.Println("你好, 世界")
			wg.Done()
		}()
	}

	// 等待N个后台线程完成
	wg.Wait()
}
```

### 注意：

如果不等待`channel`完成，或者使用`sync.waitgroup`，`main`会因为线程完成直接退出，而父进程`main`退出，会终止掉所有的子进程，会使得`channel`没有全部执行就退出程序。

