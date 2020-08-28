package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// l : *TCPListener
	l, err := net.Listen("tcp", "0.0.0.0:12345")
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println("---------------------------------------------------")
		// cLeft : *TCPConn
		cLeft, err := l.Accept()
		if err != nil {
			fmt.Println("err:", err)
			continue
		}

		go func() {
			defer cLeft.Close()
			cRight, err := net.Dial("tcp", "220.181.38.148:80") // IP of baidu.com
			defer cRight.Close()

			n1, n2, err := RelayTCP(cLeft, cRight)
			fmt.Println(n1, n2, err)
		}()

	}
}

// RelayTCP copies between left and right bidirectionally. Returns number of
// bytes copied from right to left, from left to right, and any error occurred.
func RelayTCP(left, right net.Conn) (int64, int64, error) {
	type res struct {
		N   int64
		Err error
	}
	ch := make(chan res)

	go func() {
		n, err := io.Copy(right, left)
		fmt.Println("left->right", n, err)
		// right.SetDeadline(time.Now()) // wake up the other goroutine blocking on right
		// left.SetDeadline(time.Now()) // wake up the other goroutine blocking on left
		ch <- res{n, err}
	}()

	n, err := io.Copy(left, right)
	fmt.Println("right->left", n, err)
	// right.SetDeadline(time.Now()) // wake up the other goroutine blocking on right
	// left.SetDeadline(time.Now()) // wake up the other goroutine blocking on left
	rs := <-ch

	if err == nil {
		err = rs.Err
	}
	return n, rs.N, err
}
