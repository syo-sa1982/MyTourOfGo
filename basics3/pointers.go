package main

import "fmt"

func main() {
	i, j := 42, 2701


	fmt.Println("### ポインタ定義前 ###")
	fmt.Println(&i)
	fmt.Println(&j)
	p := &i
	fmt.Println("### i:代入前 ###")
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(i)
	*p = 21
	fmt.Println("### i:代入後 ###")
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(i)

	p = &j
	fmt.Println("### j:代入前 ###")
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(j)
	*p = *p / 37
	fmt.Println("### j:代入後 ###")
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(j)
}
