package main

import "fmt"

func main() {
	var (
		h0, m0, h1, m1 int
	)

	fmt.Scanln(&h0, &m0, &h1, &m1)
	horas := h1 - h0
	minutos := m1 - m0

	if horas >= 0 && minutos > 0 {
		// Continua a mesma formatação
	} else if horas == 0 && minutos == 0 {
		horas = 24
	} else if horas > 0 && minutos < 0 {
		horas--
		minutos += 60
	} else if horas < 0 && minutos >= 0 {
		horas += 24
	} else if horas <= 0 && minutos < 0 {
		horas += 24
		horas--
		minutos += 60
	}

	fmt.Printf("O JOGO DUROU %v HORA(S) E %v MINUTO(S)\n", horas, minutos)
}
