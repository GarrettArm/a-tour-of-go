package main

import (
	"fmt"
)

type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	describe(i) // (<nil>, <nil>)
	i.M()       // panic
	// nil interface has no value or concrete type
	// i.M() cannot know which type M() belongs to.
}
