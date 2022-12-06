package main

import "fmt"

func main() {
	var N int

	fmt.Scan(&N)
	X, Y := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&X[i], &Y[i])
	}
	for key := range X {
		soma := somaImp(X[key], Y[key])
		fmt.Printf("%d\n", soma)
	}

}

func somaImp(a, b int) int {
	soma := 0
	if a < b {
		for i := a + 1; i < b; i++ {
			if i%2 != 0 {
				soma += i
			}
		}
	} else if a > b {
		for i := b + 1; i < a; i++ {
			if i%2 != 0 {
				soma += i
			}
		}
	}
	return soma
}
