package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func (t *testing.T, resultado, esperado string)  {
		t.Helper()
		if resultado != esperado {
			t.Errorf("Resultado: '%s', esperado: '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Levison", "")
		esperado := "Olá, Levison"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, Mundo"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Levison", "espanhol")
		esperado := "Hola, Levison"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em francês", func(t *testing.T){
		resultado := Ola("Levison", "francês")
		esperado := "Bonjour, Levison"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em inglês", func(t *testing.T) {
		resultado := Ola("Levison", "inglês")
		esperado := "Hello, Levison"

		verificaMensagemCorreta(t, resultado, esperado)
	})
}
