package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlices(s)

	s = s[:0]
	printSlices(s)

	s = s[:5]
	printSlices(s)

	s = s[:4]
	printSlices(s)

	s = s[:6] // 容量が6以内なら代入できる
	printSlices(s)

	// 容量超えるとエラー
	//s = s[:7]
	//printSlices(s)

	s = s[2:]
	printSlices(s)

	s = s[1:]
	printSlices(s)

	s = s[2:]
	printSlices(s)

	// nil slices
	var s0 []int
	printSlices(s0)
	if s0 == nil {
		fmt.Println("nil!")
	}
}

func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}