package main

import "fmt"

func fibonacci() func() int {
	a, b, c := 0, 0, 1
	return func() int {
		a, b, c = b, c, b+c
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
	// return fib sequence (0, 1, 1, 2, 3, 5, ...)
}
