package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	var t *T
	i = t
	describe(i) // (<nil>, *main.T)
	i.M()       // <nil>
	// empty structure is nil
	// but interface with empty value is non-nil
	// nil-receiver.M() can be handled gracefully.

	i = &T{"hello"}
	describe(i) // (&{hello}, *main.T)
	i.M()       // hello
}
