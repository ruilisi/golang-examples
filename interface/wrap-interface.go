package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type UpperWriter struct {
	io.Writer
}

func (p *UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}

func main() {
	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")
}
