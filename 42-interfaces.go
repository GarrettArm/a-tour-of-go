package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a, b Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f                // MyFloat implements Abser
	b = &v               // *Vertex implements Abser
	fmt.Println(b.Abs()) // 5
	fmt.Println(a.Abs()) // 1.4142

	// var c Abser
	// c = v // Vertex does not implement Abser // panic
	// fmt.Println(c.Abs())  // cannot work
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
