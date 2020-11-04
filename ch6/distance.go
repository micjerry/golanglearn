package main

import (
	"fmt"
	"gopl.io/ch6/geometry"
)

func main() {
	perim := geometry.Path{{1,1}, {5,1}, {5,4}, {1, 1}}
	fmt.Println(perim.Distance())
	
	r := &geometry.Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)
}