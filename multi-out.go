package main

import "fmt"

// cool omission of type & multi-return from function using commas :0
func swap(x, y string) (string, string) {
	return y,x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
