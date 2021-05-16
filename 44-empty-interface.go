package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i interface{}
	// can hold value of any type
	// because any type implements at least 0 methods
	// useful for dealing with values of unknown type.

	describe(i) // (<nil>, <nil>)

	i = 42
	describe(i) // (42, <nil>)

	i = "hello"
	describe(i) // (hello, <nil>)
}
