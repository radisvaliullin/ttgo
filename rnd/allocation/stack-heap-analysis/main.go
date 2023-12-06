package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

// go build -gcflags '-m' -o ./bin/main ./test/allocation/stack-heap-analysis
// escapes to heap - ?
// moved to head - ?
//
func main() {

	var m1, m2, m3, m4 runtime.MemStats

	runtime.ReadMemStats(&m1)
	i := 33
	i2 := 44
	i3 := 55
	runtime.ReadMemStats(&m2)
	i4 := 66
	runtime.ReadMemStats(&m3)

	go func() *int {
		return &i4
	}()

	pi := uintptr(unsafe.Pointer(&i))
	pi2 := uintptr(unsafe.Pointer(&i2))
	pi3 := uintptr(unsafe.Pointer(&i3))
	pi4 := uintptr(unsafe.Pointer(&i4))

	fmt.Printf("%x %x %x %x %v\n", pi, pi2, pi3, pi4, i4)

	runtime.ReadMemStats(&m4)

	heapStat(&m1, &m2)
	heapStat(&m2, &m3)
	heapStat(&m3, &m4)
}

func heapStat(m1, m2 *runtime.MemStats) {
	fmt.Printf("heap upped: %v\n", m2.Alloc-m1.Alloc)
}
