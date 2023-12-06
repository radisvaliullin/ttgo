package main

import "sync"

func main() {

	ch := getCh()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(ch chan obj) {
		wg.Done()
		o := <-ch
		o2 := <-ch
		o.sl = o2.sl
	}(ch)

	wg.Wait()
}

type obj struct {
	sl []byte
}

func getCh() chan obj {
	ch := make(chan obj, 1024)

	ob := obj{}
	ch <- ob
	ob2 := obj{sl: make([]byte, 1024)}
	ch <- ob2

	return ch
}
