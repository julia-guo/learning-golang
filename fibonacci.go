package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first := 0
	second := 1
	now := 0
	// interesting, returns a function closure
	return func() int {
		now = first + second
		first = second
		second = now
		return first
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

