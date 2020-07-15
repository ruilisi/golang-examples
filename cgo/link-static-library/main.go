package main

//#cgo CFLAGS: -I.
//#cgo LDFLAGS: -L${SRCDIR} -lnumber
//
//#include "number.h"
import "C"
import "fmt"

func main() {
	fmt.Println(C.number_add_mod(10, 5, 12))
}
