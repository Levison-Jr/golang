package main

import "fmt"


func main() {
	var (
		N int
	)

	fmt.Scan(&N)
	X := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&X[i])
	}

	for _, value := range X {
		switch {
		case value%2 == 0 && value > 0:
			fmt.Printf("EVEN POSITIVE\n")
		case value%2 == 0 && value < 0:
			fmt.Printf("EVEN NEGATIVE\n")
		case value%2 != 0 && value > 0:
			fmt.Printf("ODD POSITIVE\n")
		case value%2 != 0 && value < 0:
			fmt.Printf("ODD NEGATIVE\n")
		case value == 0:
			fmt.Printf("NULL\n")
		}
	}
}