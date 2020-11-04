package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(5, 6))
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}