package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola()
	esperado := "Olá, mundo"

	if resultado != esperado {
		t.Errorf("Resultado: '%s', esperado: '%s'", resultado, esperado)
	}
}
