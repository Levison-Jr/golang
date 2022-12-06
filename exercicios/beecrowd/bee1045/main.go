package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var (
		A, B, C float64
	)

	fmt.Scanln(&A, &B, &C)
	valores := []float64{A, B, C}

	sort.Slice(valores, func(i, j int) bool {
		return valores[i] > valores[j]
	})

	A, B, C = valores[0], valores[1], valores[2]

	switch {
	case A >= B+C:
		fmt.Printf("NAO FORMA TRIANGULO\n")
	case math.Pow(A, 2) == math.Pow(B, 2)+math.Pow(C, 2):
		fmt.Printf("TRIANGULO RETANGULO\n")
	case math.Pow(A, 2) > math.Pow(B, 2)+math.Pow(C, 2):
		fmt.Printf("TRIANGULO OBTUSANGULO\n")
	case math.Pow(A, 2) < math.Pow(B, 2)+math.Pow(C, 2):
		fmt.Printf("TRIANGULO ACUTANGULO\n")
	}

	if A == B && B == C && C == A {
		fmt.Printf("TRIANGULO EQUILATERO\n")
	} else if A == B || B == C || C == A {
		fmt.Printf("TRIANGULO ISOSCELES\n")
	}
}
