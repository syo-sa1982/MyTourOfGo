package main

import (
	"fmt"
	"math"
)

/*
最初は、その計算式を10回だけ繰り返し、
x を(1, 2, 3, ...)と様々な値に対する結果が
どれだけ正解値に近づくかを確認してみてください。
 */
func Sqrt(x float64) float64 {
	fmt.Println("### Start Sqrt ###")
	z := float64(1)

	for i := 0; i < 10; i++ {
		z2 := z - (z * z - x) / (2 * z)
		z = z2
	}

	return z
}

func Sqrt2(x float64) float64 {
	fmt.Println("### Start Sqrt2 ###")
	z := float64(1)

	/*
	次に、ループを回すときの直前に求めたzの値が
	これ以上変化しなくなったとき(または、差がとても小さくなったとき)に
	停止するようにループを変更してみてください。
	この変更により、ループ回数が多くなったか、少なくなったのか見てみてください。
	math.Sqrt と比べてどれくらい近似できましたか？
	 */
	for i := 0; true; i++ {
		z2 := z - (z * z - x) / (2 * z)
		// 0.000000000000001 より差が小さくならないっぽい(これより小数点の桁増やすと無限ループ)
		if z == z2 || math.Abs(z2 - z) < 0.000000000000001 {
			break
		}
		z = z2
		fmt.Println(i)
	}

	return z
}

func main() {
	x := float64(2) // change, me!

	fmt.Println(Sqrt(x))
	fmt.Println(Sqrt2(x))
	fmt.Println("### gos Sqrt ###")
	fmt.Println(math.Sqrt(x))
}