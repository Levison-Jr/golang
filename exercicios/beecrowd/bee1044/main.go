package main

import "fmt"

func main() {
	var (
		A, B int
	)

	fmt.Scanln(&A, &B)
	if A > B {
		switch {
		case A%B == 0:
			fmt.Printf("Sao Multiplos\n")
		default:
			fmt.Printf("Nao sao Multiplos\n")
		}
	} else {
		switch {
		case B%A == 0:
			fmt.Printf("Sao Multiplos\n")
		default:
			fmt.Printf("Nao sao Multiplos\n")
		}
	}
}
