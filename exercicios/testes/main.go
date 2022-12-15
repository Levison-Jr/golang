package main

import (
	"fmt"
)

func main() {
	var (
		opt, alcool, gasolina, diesel int
	)

	for {
		fmt.Scan(&opt)
		if opt == 4 {
			break
		}
		switch opt {
		case 1:
			alcool++
		case 2:
			gasolina++
		case 3:
			diesel++
		}
	}
	fmt.Println("MUITO OBRIGADO")
	fmt.Printf("Alcool: %d\n", alcool)
	fmt.Printf("Gasolina: %d\n", gasolina)
	fmt.Printf("Diesel: %d\n", diesel)
}
