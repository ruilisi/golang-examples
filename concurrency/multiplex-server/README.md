## multiplex-server

### 运行方式

`go run main.go`


### 过程

1.新建一个存放请求的容器

`adder, quit := startServer(func(a, b int) int { return a + b })` 

2.模拟接受到请求,向容器里面添加请求


3.对于容器中的数据执行特定的 `op` 函数

```
func run(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b)
}
```

4.循环判断每个请求的值是否正确
因为每一个请求都是独立的，所以对于每一个请求,判断的顺序无关紧要

```
	for i := N - 1; i >= 0; i-- { // doesn't matter what order
		if <-reqs[i].replyc != N+2*i {
			fmt.Println("fail at", i)
		} else {
			fmt.Println("Request ", i, " is ok!")
		}
	}

```




