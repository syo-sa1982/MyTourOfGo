package main

import "fmt"

type Vertex3 struct {
	Lat, Long float64
}

var m map[string]Vertex3

var m3 = map[string]Vertex3{
	"Bell Labs": {40.68433, -74.39967},
	"Google": {37.42202, -122.08408},
}

func main() {
	m = make(map[string]Vertex3)
	m["Bell Labs"] = Vertex3{
		40.68433, -74.39967,
	}

	var m2 = map[string]Vertex3{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m2)
	fmt.Println(m3)
}
