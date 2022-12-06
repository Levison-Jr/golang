package main

import "fmt"

func main() {
	var (
		h0, h1 int
	)

	fmt.Scanln(&h0, &h1)
	horas := h1 - h0

	if horas > 0 {
		// Continua a mesma formatação
	} else if horas == 0 {
		horas = 24
	} else if horas < 0 {
		horas += 24
	}

	fmt.Printf("O JOGO DUROU %v HORA(S)\n", horas)
}
