package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {
	a := make([]int, 5)
	printSlice("a", a) // len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	printSlice("b", b) // len=0 cap=5 []

	c := b[:2]
	printSlice("c", c) // len=2 cap=5 [0 0]

	d := c[2:5]
	printSlice("d", d) // len=3 cap=3 [0 0 0]
}
