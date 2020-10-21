# c语言调用golang库

### 生成动态库的命令

`go build  -buildmode=c-shared -o libhello.so   main.go`

### 生成静态库的命令

` go build -buildmode=c-archive -o libhello.a main.go`

### 编译命令

`gcc main.c -o main -I./ -L./ -lsum -lpthread`



### 运行方法

`make all`

`./main`

`make clean`