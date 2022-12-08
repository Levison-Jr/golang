package inteiros

import (
	"fmt"
	"testing"
)

func TestSoma(t *testing.T) {
	soma := Soma(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("esperado: '%d', resultado: '%d'", esperado, soma)
	}
}

func ExampleSoma() {
	soma := Soma(1, 5)
	fmt.Println(soma)
	// Output: 6
}
