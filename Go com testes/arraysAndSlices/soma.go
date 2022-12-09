package main

func Soma(numeros []int) int {
	soma := 0
	for _, value := range numeros {
		soma += value
	}
	return soma
}
