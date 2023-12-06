package main

import "fmt"

// go build -gcflags '-m' -o ./bin/main ./test/allocation/stack-heap-asm
// escapes to heap - ?
// moved to head - ?
//
// go tool compile -S ./test/allocation/stack-heap-asm/main.go > ./test/allocation/stack-heap-asm/asm.out
// this asm output do not clear situation about allocation.
func main() {
	i := 33
	i2 := 66
	fmt.Println(i, &i2)
	i3 := i + i2 + 2
	_ = i3
}
