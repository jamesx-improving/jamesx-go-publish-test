package api

// #cgo LDFLAGS: -ljamesx_go_publish_test
// #cgo LDFLAGS: -L${SRCDIR}/../build
// #include "../lib.h"
import "C"

// This function prints hello world from rust
func GoWithRustPrint() {
	C.print_hello()
}

// func main() {
// 	GoWithRustPrint()
// }

// #include "../lib.h"
// void print_hello();
