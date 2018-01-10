package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// Struct Literals
type Vertex2 struct {
	X, Y int
}
var (
	v1 = Vertex2{1, 2}
	v2 = Vertex2{X: 1}
	v3 = Vertex2{}
	p1 = &Vertex2{1, 2}
)

func main() {
	fmt.Println(Vertex{1, 2})

	// Struct Fields
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// Pointers to structs
	p := &v
	fmt.Println(v)
	fmt.Println(p)
	fmt.Println(*p)
	p.X = 1e9
	fmt.Println(v)
	fmt.Println(p)
	fmt.Println(*p)

	// Struct Literals
	fmt.Println(v1, p1, *p1, v2, v3)
}
