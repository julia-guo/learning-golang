package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// methods live outside structs
// emit infinite stream of 'A'
func (m MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
