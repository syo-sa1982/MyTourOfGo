package main

import "fmt"

func main() {
	i := 42
	s := "42"
	f := 3.142
	c := 0.867 + 0.5i

	fmt.Printf("This is of type %T\n", i)
	fmt.Printf("This is of type %T\n", s)
	fmt.Printf("This is of type %T\n", f)
	fmt.Printf("This is of type %T\n", c)
}