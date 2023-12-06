package main

import (
	"fmt"
	"runtime"
)

func main() {

	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(1))
	runtime.GC()

	// slice allocation with make
	var startUsingMake, endUsingMake runtime.MemStats
	runtime.ReadMemStats(&startUsingMake)
	sl := make([]byte, 8)
	l := len(sl)
	// go func() { fmt.Println(len(sl)) }()
	runtime.ReadMemStats(&endUsingMake)
	allocUsingMake := endUsingMake.TotalAlloc - startUsingMake.TotalAlloc
	fmt.Printf("sl len %v, alloc:%d \n", l, allocUsingMake)
	// fmt.Printf("sl %v\n", sl)

	// array allocation
	var startUsingArr, endUsingArr runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&startUsingArr)
	var arr [8]byte
	// go func() { fmt.Println(len(arrUsingVar)) }()
	runtime.ReadMemStats(&endUsingArr)
	allocArr := endUsingArr.TotalAlloc - startUsingArr.TotalAlloc
	fmt.Printf("arr %v, alloc:%d \n", arr, allocArr)
	fmt.Printf("arr %v\n", arr)
}
