package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go func(){
    Walk(t1, c1)
    close(c1)
  }()
	go func() {
    Walk(t2, c2)
    close(c2)
  }()

  arr := []int {}

  for x := range c1 {
    arr = append(arr, x)
  }
  same := true
  index := 0
  for x := range c2 {
    if (index >= len(arr)) {
      same = false
    }
    if arr[index] != x {
      same = false
    }
    index += 1
  }

	return same
}

func main() {
	fmt.Printf("same: %v\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("same: %v\n", Same(tree.New(1), tree.New(2)))
}
