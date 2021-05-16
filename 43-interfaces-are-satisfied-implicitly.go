package main

import (
	"fmt"
	// "math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

type R struct {
	S string
}

func (r *R) M() {
	if r == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(r.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I

	// i = &T{"Hello"}
	// describe(i) // (&{Hello}, *main.T)
	// i.M()       // Hello

	// i = F(math.Pi)
	// describe(i) // (3.1415, main.F)
	// i.M()       // 3.1415

	// var r *R
	// i = r
	// describe(i) // (<nil>, *main.R)
	// i.M()       // <nil>

	describe(i) // (<nil>, <nil>)
	i.M()       // run-time error
	// nil interface has no value or concrete type
	// i.M() cannot know which type M() belongs to.
}
