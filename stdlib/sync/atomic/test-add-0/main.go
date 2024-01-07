package main

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
)

var inrem int64

var intMap = map[int64]struct{}{}

var numGo = 10_000
var intChan = make(chan int64, numGo)

var doneChan = make(chan struct{})

func main() {
	// Let say multiple goroutine tries increment some value using atomic add
	// is there chance that they get same return value
	// it should not but I am paranoid
	// little be investigated about atomic (c/c++) and looked source code it should't

	go func() {

		wg := sync.WaitGroup{}

		for i := 0; i < numGo; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					inc := atomic.AddInt64(&inrem, 1)
					if inc < 0 {
						return
					}
					intChan <- inc
				}
			}()
		}

		wg.Wait()
		doneChan <- struct{}{}
	}()

	loopDone := false
	cnt := 0
	for {
		if loopDone {
			break
		}
		select {
		case <-doneChan:
			loopDone = true
		case i := <-intChan:
			if i%1000_000 == 0 {
				cnt++
				fmt.Println("NEXT LOOP LIMIT", cnt)
			}
			if _, ok := intMap[i]; ok {
				fmt.Printf("FOUND duplicate %v\n", i)
				os.Exit(42)
			} else {
				intMap[i] = struct{}{}
			}
		}
	}
}
