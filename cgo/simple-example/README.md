## go 调用c的基本方式

```
/*
  这边放c的代码,包括c的头文件
*/
import "C"  // 一定要有这个include，并且一定要在c代码下面一行

func main (){
  // 通过C.func()调用注释中的c函数
  C.func()

}

```

### 运行方式

`go run 文件名`

