package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		N, contFood        int
		linhaTab, tabTotal string
		foods              []int
	)

	fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		fmt.Scan(&linhaTab)
		if len(linhaTab) != N {
			fmt.Printf("Você digitou um valor diferente de %v, favor digitar a linha %v novamente...\n", N, i)
			i--
			continue
		} else if i%2 == 0 {
			linhaTab = inv(linhaTab)
		}
		tabTotal += linhaTab
	}

	sliceTab := []byte(tabTotal)

	for _, value := range sliceTab {
		if string(value) == "o" {
			contFood++
			foods = append(foods, contFood)
		}
		if string(value) == "A" {
			contFood = 0
		}
	}
	sort.Slice(foods, func(i, j int) bool {
		return foods[i] > foods[j]
	})
	fmt.Println(foods[0])
}

func inv(s string) string {
	y := []byte(s)

	for i := 0; i < len(s); i++ {
		y[i] = s[len(s)-1-i]
	}
	return string(y)
}
