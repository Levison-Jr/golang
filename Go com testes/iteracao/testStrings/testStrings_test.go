package teststrings

import (
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	contagem := strings.Count("oooOlLaaa", "o")
	esperado := 3
	if contagem != esperado {
		t.Errorf("contagem: '%d', esperado: '%d'", contagem, esperado)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Count("23871112387462871143948398411192837139187211", "1")
	}
}
