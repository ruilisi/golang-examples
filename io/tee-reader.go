package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := io.TeeReader(strings.NewReader("Go语言中文网"), os.Stdout)
	buf := make([]byte, 20)
	n, _ := reader.Read(buf)
	fmt.Println(string(buf[:n]))
}
