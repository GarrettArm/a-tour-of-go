package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println(m["Answer"]) // 42

	m["Answer"] = 48
	fmt.Println(m["Answer"]) // 48

	delete(m, "Answer")
	fmt.Println(m["Answer"]) // 0

	v, ok := m["Answer"]
	fmt.Println(v, ok) // 0 false
}
