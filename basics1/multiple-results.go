package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("一生世界に", "挨拶してろ")
	fmt.Println(a, b)
}