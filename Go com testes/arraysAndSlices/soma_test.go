package main

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {

	numeros := []int{1, 2, 3, 4, 5}

	resultado := Soma(numeros)
	esperado := 15

	if resultado != esperado {
		t.Errorf("resultado: '%d', esperado: '%d', dado: '%v'", resultado, esperado, numeros)
	}
}

func TestSomaAll(t *testing.T) {

	resultado := SomaAll([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("resultado: %d, esperado: %d", resultado, esperado)
	}
}
