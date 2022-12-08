package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	resultado := Repetir("b", 5)
	esperado := "bbbbb"

	if resultado != esperado {
		t.Errorf("esperado: '%s', resultado: '%s'", esperado, resultado)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 7)
	}
}

func ExampleRepetir() {
	repetir := Repetir("a", 5)
	fmt.Println(repetir)
	// Output: aaaaa
}
