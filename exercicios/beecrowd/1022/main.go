package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var (
		N, N1, D1, N2, D2, num, den int
		operador                    string
	)

	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d / %d %s %d / %d", &N1, &D1, &operador, &N2, &D2)

		if operador == "+" {
			num = N1*D2 + N2*D1
			den = D1 * D2
		} else if operador == "-" {
			num = N1*D2 - N2*D1
			den = D1 * D2
		} else if operador == "*" {
			num = N1 * N2
			den = D1 * D2
		} else if operador == "/" {
			num = N1 * D2
			den = N2 * D1
		}
		mdc := mdc(num, den)
		if mdc > 0 {
			numRed := num / mdc
			denRed := den / mdc
			fmt.Printf("%d/%d = %d/%d\n", num, den, numRed, denRed)
		} else {
			fmt.Printf("%d/%d = %d/%d\n", num, den, num, den)
		}
	}
}

func mdc(x, y int) int {
	num := []int{int(math.Abs(float64(x))), int(math.Abs(float64(y)))}
	sort.Ints(num)

	for i := num[0]; i > 1; i-- {
		if x%i == 0 && y%i == 0 {
			return i
		}
	}

	return -1
}
