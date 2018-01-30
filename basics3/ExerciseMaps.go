package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	wordSlice := strings.Fields(s)

	for i := range wordSlice {
		m[wordSlice[i]]++
	}
	return m
}


func main()  {
	wc.Test(WordCount)
}
