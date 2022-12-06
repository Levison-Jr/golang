package main

import "fmt"

func main() {
	var (
		salario, reajuste, reajusteP float64
		percent                      string = "%"
	)

	fmt.Scan(&salario)
	switch {
	case salario >= 0 && salario <= 400:
		reajusteP = 0.15
		reajuste = salario * reajusteP
		salario += reajuste
	case salario >= 400.01 && salario <= 800:
		reajusteP = 0.12
		reajuste = salario * reajusteP
		salario += reajuste
	case salario >= 800.01 && salario <= 1200:
		reajusteP = 0.10
		reajuste = salario * reajusteP
		salario += reajuste
	case salario >= 1200.01 && salario <= 2000:
		reajusteP = 0.07
		reajuste = salario * reajusteP
		salario += reajuste
	case salario > 2000:
		reajusteP = 0.04
		reajuste = salario * reajusteP
		salario += reajuste
	}

	fmt.Printf("Novo salario: %.2f\n", salario)
	fmt.Printf("Reajuste ganho: %.2f\n", reajuste)
	fmt.Printf("Em percentual: %.0f %v\n", reajusteP*100, percent)
}
