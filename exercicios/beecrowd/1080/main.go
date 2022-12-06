package main

import (
	"fmt"
	"sort"
)

func main() {
	valores := make([]int, 100)
	x := make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		fmt.Scan(&valores[i])
	}

	x = append(x, valores...)
	sort.Slice(x, func(i, j int) bool {
		return x[i] > x[j]
	})

	valor, index := maior(valores, x[0])
	if valor != -1 || index != -1 {
		fmt.Printf("%d\n", valor)
		fmt.Printf("%d\n", index+1)
	} else {
		fmt.Println("ERRO!")
	}
}

func maior(si []int, num int) (int, int) {
	for i, v := range si {
		if v == num {
			return v, i
		}
	}

	return -1, -1
}
