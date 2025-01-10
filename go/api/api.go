package api

// #cgo LDFLAGS: -L../target/release -ljamesx_go_publish_test
//
// void print_hello();
import "C"

func GoWithRustPrint() {
	C.print_hello()
}

// func main() {
// 	GoWithRustPrint()
// }

// #include "../lib.h"
