package main

import "fmt"

// continued
func add(x, y int) int {
	return x + y
}



func main() {
	hoge := add(42, 13)
	fmt.Println(hoge)
	fmt.Println(add(42, 13))
}
