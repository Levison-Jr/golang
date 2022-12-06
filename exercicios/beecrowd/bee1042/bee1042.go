package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		n1, n2, n3 int
	)

	fmt.Scanln(&n1, &n2, &n3)
	valores := []int{n1, n2, n3}
	sort.Ints(valores)

	for _, value := range valores {
		fmt.Printf("%v\n", value)
	}

	fmt.Printf("\n%v\n%v\n%v\n", n1, n2, n3)
}
