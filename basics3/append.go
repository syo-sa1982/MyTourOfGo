package main

import "fmt"

func main() {
	var s []int
	printSlice2(s)
}

func printSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}