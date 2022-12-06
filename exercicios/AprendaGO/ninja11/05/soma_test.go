package main

import "testing"

func TestSoma(t *testing.T) {
	v := soma(10, 20)

	if v != 30 {
		t.Errorf("Erro! Esperado: 30 || Obtido: %v", v)
	}
}
