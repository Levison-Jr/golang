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

func SomaTodoOResto(numeros ...[]int) []int {
	var total []int
	for _, value := range numeros {
		if len(value) == 0 {
			total = append(total, 0)
		} else {
			final := value[1:]
			total = append(total, Soma(final))
		}
	}

	return total
}
