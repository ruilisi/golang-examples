package main

import (
	"fmt"
	"testing"
)

type TB struct {
	testing.TB
}

func (p *TB) Fatal(args ...interface{}) {
	fmt.Println("[TB.Fatal is replaced]", args)
}

func main() {
	var tb testing.TB = new(TB)
	tb.Fatal("Hello, playground")
}
