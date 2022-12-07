package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Levison")
	esperado := "Olá, Levison"

	if resultado != esperado {
		t.Errorf("Resultado: '%s', esperado: '%s'", resultado, esperado)
	}
}
