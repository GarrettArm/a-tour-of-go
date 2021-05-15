package main

import "fmt"

/*
  when a slice grows beyond its underlying array,
  it switches to use a new, larger underlying array
*/

func main() {
	var s []int
	printSlice(s) // len=0 cap=0 []

	s = append(s, 0)
	printSlice(s) // len=1 cap=1 [0]

	s = append(s, 2, 3, 4)
	printSlice(s) // len=4 cap=4 [0 2 3 4]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
