package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		N, num int
	)

	fmt.Scanln(&N)
	valores := make([]int, 0, N)

	for i := 0; i < N; i++ {
		fmt.Scanln(&num)
		valores = append(valores, num)
	}

	for _, value := range valores {
		testePrimo := Primo(value)

		if testePrimo {
			fmt.Printf("Prime\n")
		} else {
			fmt.Printf("Not Prime\n")
		}
	}
}

func Primo(p int) bool {
	for j := 2; j <= int(math.Sqrt(float64(p))); j++ {
		if p%j == 0 {
			return false
		}
	}
	return true
}
