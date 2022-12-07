package main

import "fmt"

func main() {
	var (
		notas        []float64 = make([]float64, 0, 2)
		valor, media float64
		novo         int
	)

	for {
		for len(notas) < 2 {
			fmt.Scan(&valor)
			if valor >= 0 && valor <= 10 {
				notas = append(notas, valor)
			} else {
				fmt.Println("nota invalida")
			}
		}

		media = (notas[0] + notas[1]) / 2
		fmt.Printf("media = %.2f\n", media)
		for {
			fmt.Println("novo calculo (1-sim 2-nao)")
			fmt.Scan(&novo)
			if novo == 1 || novo == 2 {
				notas = make([]float64, 0, 2)
				break
			}
		}
		if novo == 2 {
			break
		}
	}
}
