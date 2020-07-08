package packageB

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("Before entering main func, go runs in a single goroutine")
	time.Sleep(time.Second * 2)
	fmt.Println("init packageB b")
}
