package main

import "fmt"

func main() {
	/*
		A pointer holds a memory address of a value.
	*/
	i, j := 42, 2701

	p := &i               // point to i
	fmt.Printf("%T\n", p) // var p is type *int
	fmt.Println(*p)       // read i through the pointer.  *p is 42
	*p = 21               // set  i through the pointer
	fmt.Println(i, *p)    // i and *p are 21

	p = &j             // point to j
	*p = *p / 37       // divide j through the pointer
	fmt.Println(j, *p) // j and *p are 73

	// & creates a pointer
	// * defererences, indirects a pointer
}
