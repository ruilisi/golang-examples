## invoke-c

### 运行方式

`go build .`
`./invoke-c`

[main.go](main.go) 中调用 `C.SayHello(C.CString("Hello, World\n"))`

`C.SayHello()` 实际定义在 `hello.c` 中

因而 `main.go` 中调用的 `C.SayHello()` 是 `hello.c` 中的函数

