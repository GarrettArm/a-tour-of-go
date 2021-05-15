package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // {1 2}
	v2 = Vertex{X: 1}  // {1 0}  default values for unspecifieds
	v3 = Vertex{}      // {0 0}
	p1 = &Vertex{3, 4} // has type *Vertex.  The Vertex itself is anonymous
)

func main() {
	v := Vertex{5, 6}
	p2 := &v
	p2.X = 1e9
	// (*p).X would be most accurate, but go gives this syntatic shortcut

	fmt.Println(v)              // {1e9 6}
	fmt.Println(v1, v2, v3, p1) // {1,2} {1 0} {0 0} &{3 4}
}
