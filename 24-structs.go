package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	// (*p).X would be most accurate, but go gives this syntatic shortcut
	fmt.Println(v)
}
