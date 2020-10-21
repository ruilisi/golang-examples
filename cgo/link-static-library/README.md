# golang调用c语言库

### 编译命令(生成c静态库)

`gcc -c -o main.o main.c`

`ar rcs libmain.a main.o`



### 运行方式

`make`

`go run main.go`

`make clean`