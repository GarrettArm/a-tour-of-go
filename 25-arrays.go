package main

import "fmt"

/*
  Arrays cannot be resized.
  And they contain members of the same type
*/

func main() {
	var a [2]string // a is array of 2 strings
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
