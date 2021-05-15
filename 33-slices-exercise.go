package main

import "golang.org/x/tour/pic"

/*
  Only works in Golang.org/tour platform
*/

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			row[x] = uint8(y ^ x)
		}
		image[y] = row
	}
}

func main() {
	pic.Show(Pic(16, 16))
}
