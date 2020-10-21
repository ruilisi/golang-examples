package main

/*
#include <stdio.h>
void test() {
	printf("hello world\n");
}
*/
import "C"

func main() {
	C.test()
}
