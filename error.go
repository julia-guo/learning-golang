package main

import (
	"fmt"
)

// error type
type ErrNegativeSqrt float64

// error catcher function
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// return error function if bad input
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := float64(x/2)
	delta := float64((z*z - x) / (2*z))

	for delta > 0.000001 {
		delta = (z*z - x) / (2*z)
		z -= delta
	}
	return z, nil
}

// will throw error when -2 inputted
func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

