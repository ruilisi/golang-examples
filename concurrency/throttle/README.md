## throttle 限流

### 1. 运行过程

1. 模拟发送请求

```
	go func() {
		for {
			requests <- "New request"
		}
	}()
```

2. 控制接受请求时长

```
  chRate := time.Tick(time.Second)
```

3. 依次处理每个请求

```
	for req := range requests {
		<-chRate
		fmt.Println(req)
	}
```
### 2. 结果

每隔1秒输出一个请求
