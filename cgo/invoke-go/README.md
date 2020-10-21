## invoke-go

### 运行方式

`go build .`
`./invoke-go`

[main.go](main.go) 中调用 `C.SayHello(C.CString("Hello, World\n"))`

`C.SayHello()` 实际定义在 `hello.go` 中

因而 `main.go` 中调用的 `C.SayHello()` 是 `hello.go` 中的函数
