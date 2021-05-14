package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Printf("%v is a %T\n", x, x)
	fmt.Printf("%v is a %T\n", f, f)
	fmt.Printf("%v is a %T\n", z, z)
}
