package main

import "fmt"

func main() {
	fmt.Println(Ola("mundo", ""))
}

const espanhol = "espanhol"
const prefixoOlaPortugues = "Olá, "
const prefixoHolaEspanhol = "Hola, "

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}
	if idioma == espanhol {
		return prefixoHolaEspanhol + nome
	}
	return prefixoOlaPortugues + nome
}
