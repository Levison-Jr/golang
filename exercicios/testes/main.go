package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		X, Y int
		soma int = 0
	)

	fmt.Scanln(&X, &Y)
	input := []int{X, Y}
	sort.Ints(input)

	for i := input[0]; i <= input[1]; i++ {
		if i%13 != 0 {
			soma += i
		}
	}
	fmt.Printf("%d\n", soma)
}
