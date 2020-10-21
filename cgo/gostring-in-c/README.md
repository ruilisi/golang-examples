# //export 函数名 使用

## 运行方式

`go run main.go`

[main.go](main.go)

在调用c库的时候，可以不使用c代码，通过在go代码上面一行中中加入  `//export 函数名` 直接使用go代码（不可以含有空行）

