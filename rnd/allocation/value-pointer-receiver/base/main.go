package main

import (
	"fmt"
	"time"
)

type Obj struct {
	one int
}

func (o Obj) valueReciver() int {
	o.one = 1234
	return o.one
}

func (o *Obj) pointerReceiver() int {
	o.one = 2345
	return o.one
}

// go build -gcflags '-m' your_package
func main() {

	temp := new(int)
	_ = temp

	o := Obj{}
	go func(o Obj) {
		temp := o
		fmt.Println("goroutine: ", temp)
	}(o)

	ov := o.valueReciver()
	op := o.pointerReceiver()
	fmt.Println(ov)
	fmt.Println(op)

	time.Sleep(time.Second * 4)
}
