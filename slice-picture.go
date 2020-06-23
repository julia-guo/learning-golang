package main

import "golang.org/x/tour/pic"

// cool exercise in slicing
func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	// iterate through rows
	for i := range picture {
		picture[i] = make([]uint8, dx)
		// iterate through columns and color
		for j := range picture[i] {
			picture[i][j] = uint8((i + j) / 2)
		}
	}
	return picture
}

// paint a pretty picture
func main() {
	pic.Show(Pic)
}

