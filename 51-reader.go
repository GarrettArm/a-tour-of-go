package main

import (
	"fmt"
	"io"
	"strings"
)

/*
  io.Reader interface is used by lots of go libraries.
  it has a .Read() method
  strings.NewReader() is one library that's satisfies io.Reader interface
  byte-by-byte chunks are read, and io.EOF error marks when stream ends.
*/

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
