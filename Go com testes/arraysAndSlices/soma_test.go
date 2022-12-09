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

func TestSomaTodoOResto(t *testing.T) {

	checarResultado := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado: %d, esperado: %d", resultado, esperado)
		}
	}

	t.Run("soma numeros de cada slice passado", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}
		checarResultado(t, resultado, esperado)
	})

	t.Run("trata um slice vazio seguramente", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{}, []int{3, 4, 5})
		esperado := []int{0, 9}
		checarResultado(t, resultado, esperado)
	})

}
