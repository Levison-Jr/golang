package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		A, B, C float64
	)

	fmt.Scanln(&A, &B, &C)
	valores := []float64{A, B, C}
	sort.Float64s(valores)

	if valores[0]+valores[1] > valores[2] { // Caso seja um triângulo
		perimetro := valores[0] + valores[1] + valores[2]
		fmt.Printf("Perimetro = %.1f\n", perimetro)
	} else { // Caso não seja, logo calcula-se a área de um trapezio
		area := (A + B) * C / 2
		fmt.Printf("Area = %.1f\n", area)
	}
}
