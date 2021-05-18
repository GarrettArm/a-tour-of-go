package main

import "fmt"

type MyReader struct{}

func (m MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	r := MyReader{}

	b := make([]byte, 8)
	for {
		b, err := r.Read(b)
		fmt.Println(string(b))
		if err != nil {
			break
		}
	}
}
