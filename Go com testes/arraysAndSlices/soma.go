package main

func Soma(numeros []int) int {
	soma := 0
	for _, value := range numeros {
		soma += value
	}
	return soma
}

func SomaAll(numerosParaSoma ...[]int) []int {
	var total []int

	for _, value := range numerosParaSoma {
		total = append(total, Soma(value))
	}

	return total
}
