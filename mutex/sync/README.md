

## sync

我们创建了`setup`线程，用于对字符串`a`的初始化工作，初始化完成之后设置`done`标志为`true`。`main`函数所在的主线程中，通过`for !done {}`检测`done`变为`true`时，认为字符串初始化工作完成，然后进行字符串的打印工作。

但是Go语言并不保证在`main`函数中观测到的对`done`的写入操作发生在对字符串`a`的写入的操作之后，因此程序很可能打印一个空字符串。更糟糕的是，因为两个线程之间没有同步事件，`setup`线程对`done`的写入操作甚至无法被`main`线程看到，`main`函数有可能陷入死循环中。





根据Go语言规范，`main`函数退出时程序结束，不会等待任何后台线程。因为Goroutine的执行和`main`函数的返回事件是并发的，谁都有可能先发生，所以什么时候打印，能否打印都是未知的。

### 使用锁实现同步

``` go
func main() {
    var mu sync.Mutex

    mu.Lock()
    go func(){
        fmt.Println("你好, 世界")
        mu.Unlock()
    }()

    mu.Lock()
}
```



### 使用channel实现同步

```go
func main() {
    done := make(chan int)

    go func(){
        fmt.Println("你好, 世界")
        <-done
    }()

    done <- 1
}
```



