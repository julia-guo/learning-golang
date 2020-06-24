package main

import (
	"golang.org/x/tour/wc"
	//"fmt"
	"strings"
)

// function to count words in a string
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	//fmt.Println(words)
	m := make(map[string]int)
	for _, word := range words {
		m[word] = m[word] + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
