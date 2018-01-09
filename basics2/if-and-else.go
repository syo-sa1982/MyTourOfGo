package main

import (
	"math"
	"fmt"
)

func powElse(x, n ,lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", n, lim)
	}
	return lim
}

func main() {
	// MEMO:fmt.Printlnは引数内の処理が先に評価される
	fmt.Println(
		powElse(3, 2, 5),
		powElse(3, 2, 10),
		powElse(3, 3, 20),
		powElse(3, 3, 30),
	)
}
