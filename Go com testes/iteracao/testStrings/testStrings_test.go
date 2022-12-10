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

func TestReplaceAll(t *testing.T) {
	verificarResultado := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado: '%s', esperado: '%s'", resultado, esperado)
		}
	}

	t.Run("'old' e 'new' não vazios", func(t *testing.T) {
		resultado := strings.ReplaceAll("oink oink oink", "oink", "moo")
		esperado := "moo moo moo"
		verificarResultado(t, resultado, esperado)
	})

	t.Run("'old' vazio", func(t *testing.T) {
		resultado := strings.ReplaceAll("oink oink oink", "", "moo")
		esperado := "mooomooimoonmookmoo mooomooimoonmookmoo mooomooimoonmookmoo"
		verificarResultado(t, resultado, esperado)
	})

	t.Run("'new' vazio", func(t *testing.T) {
		resultado := strings.ReplaceAll("oink oink oink", "oink", "")
		esperado := "  "
		verificarResultado(t, resultado, esperado)
	})

}
