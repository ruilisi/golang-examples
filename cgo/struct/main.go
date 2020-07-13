package main

/*
struct A {
    int i;
    int type;
    float f;
};
*/
import "C"
import "fmt"

func main() {
	var a C.struct_A
	fmt.Println(a.i)
	fmt.Println(a.f)
	// 如果结构体的成员名字中碰巧是Go语言的关键字，可以通过在成员名开头添加下划线来访问：
	fmt.Println(a._type)
}
