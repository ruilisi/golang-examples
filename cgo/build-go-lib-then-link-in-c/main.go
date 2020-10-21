package main

//int sum(int a, int b);
//void showPrimes(int n);
import "C"

import (
	"fmt"
	"time"
)

// 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 管道过滤器: 删除能被素数整除的数
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

//export sum
func sum(a, b C.int) C.int {
	return a + b
}

//export showPrimes
func showPrimes(_n C.int) {
	n := int(_n)
	ch := GenerateNatural() // 自然数序列: 2, 3, 4, ...
	time.Sleep(1 * time.Second)
	for i := 0; i < n; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime) // 基于新素数构造的过滤器
	}
}

func main() {}
