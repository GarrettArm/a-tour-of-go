package main

import "fmt"

/*
  Arrays cannot be resized.
  And they contain members of the same type
  Slices are a dynamically-sized view into an underlying array
*/

func main() {
	var a [2]string // "a" is array of 2 strings
	a[0] = "Hello"  // manually building the array
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} // array literal
	fmt.Println(primes)

	var s []int = primes[1:4] // slice of array "a"
	fmt.Println(s)

	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)
	c := names[0:2]
	d := names[1:3]
	fmt.Println(c, d)
	d[0] = "XXX" // altering a slice alters the underlying array
	fmt.Println(c, d)
	fmt.Println(names)
}
