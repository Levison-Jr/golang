package main

import "fmt"

func main() {
	fmt.Println(Ola("mundo", ""))
}

const (
	espanhol            = "espanhol"
	frances             = "francês"
	ingles              = "inglês"
	prefixoOlaPortugues = "Olá, "
	prefixoOlaEspanhol  = "Hola, "
	prefixoOlaFrances   = "Bonjour, "
	prefixoOlaIngles    = "Hello, "
)

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixoDeSaudacao(idioma) + nome
}

func prefixoDeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case frances:
		prefixo = prefixoOlaFrances
	case ingles:
		prefixo = prefixoOlaIngles
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}
