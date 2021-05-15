package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; math.Abs(z*z-x) > 0.000005; i++ {
		fmt.Printf("Step #%v:\t%v\t%v\n", i, z, math.Abs(z*z-x))
		z -= (z*z - x) / (2 * x)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2.0))
}
