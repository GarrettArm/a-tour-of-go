package main

import "fmt"

func main() {
	defer fmt.Println("world")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("hello")
	// defers unwrap in last-in first-out
}
