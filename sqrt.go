package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(x/2)
	delta := float64((z*z - x) / (2*z))
	for delta > 0.000001 {
		delta = (z*z - x) / (2*z)
		z -= delta
	}
	return z
}

func main() {
	fmt.Println(Sqrt(7))
}
