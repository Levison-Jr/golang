package main

import "fmt"

var rato, sapo, coelho, total int
var percent string = "%"

func main() {
	var (
		N int
	)

	fmt.Scan(&N)
	quantia := make([]int, N)
	tipo := make([]string, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d %s", &quantia[i], &tipo[i])
	}

	for key, value := range tipo {
		total += quantia[key]
		switch value {
		case "R": // Rato
			rato += quantia[key]
		case "S": // Sapo
			sapo += quantia[key]
		case "C": // Coelho
			coelho += quantia[key]
		}
	}

	fmt.Printf("Total: %d cobaias\n", total)
	fmt.Printf("Total de coelhos: %d\n", coelho)
	fmt.Printf("Total de ratos: %d\n", rato)
	fmt.Printf("Total de sapos: %d\n", sapo)
	fmt.Printf("Percentual de coelhos: %.2f %s\n", (float64(coelho)/float64(total))*100, percent)
	fmt.Printf("Percentual de ratos: %.2f %s\n", (float64(rato)/float64(total))*100, percent)
	fmt.Printf("Percentual de sapos: %.2f %s\n", (float64(sapo)/float64(total))*100, percent)
}
