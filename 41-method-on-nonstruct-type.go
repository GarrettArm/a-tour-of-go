package main

import (
	"fmt"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-4)
	fmt.Println(f.Abs())
	f = MyFloat(1.24)
	fmt.Println(f.Abs())
}

/*
  Cannot only attach a method on a locally-defined type
*/
