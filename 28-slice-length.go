package main

import "fmt"

func printSlice(s []int) {
	// len(s) is length of slice
	// cap(s) is length of underlying array, starting from slice start
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	s = s[:0]
	printSlice(s) // len=0 cap=6 []
	if s == nil {
		fmt.Println("s is nil") // not a nil slice
	}

	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	s = s[2:]
	printSlice(s) // len=2 cap=6 [5 7]

	/*
		// not possible since cap=6
		s = s[2:10]
		printSlice(s)
	*/

	var n []int
	printSlice(n) // len=0 cap=0 []
	if n == nil {
		fmt.Println("n is nil") // default value of slice is nil
	}
}
