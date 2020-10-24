## mutex

### [sum.go](sum.go)

`sync.WaitGroup` 通过等待锁的机制，等待各项线程运行完毕，实现同步

一般情况下，原子操作都是通过互斥访问来保证的，通常由特殊的CPU指令提供保护。当然，如果仅仅是想模拟下粗粒度的原子操作，我们可以借助于`sync.Mute`

```go
func worker(wg *sync.WaitGroup) {
    defer wg.Done()

    for i := 0; i <= 100; i++ {
        total.Lock()
        total.value += i
        total.Unlock()
    }
}
```

在worker的循环中，为了保证`total.value += i`的原子性，我们通过`sync.Mutex`加锁和解锁来保证该语句在同一时刻只被一个线程访问。对于多线程模型的程序而言，进出临界区前后进行加锁和解锁都是必须的。如果没有锁的保护，total的最终值将由于多线程之间的竞争而可能会不正确。




