package packageB

import "fmt"

func init() {
	fmt.Println("init packageB c")
	go func() {
		fmt.Println("goroutine in init func will only run after entered main func")
	}()
}
