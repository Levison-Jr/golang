package escrevendo

import "fmt"

// X é somente um número de demonstração ;)
const X = 210

// Doc é a primeira func que ta aqui só para testes...
func Doc() {
	fmt.Println("Faz nada, só testando...")
}

// doc2 nem vai ser vista fora do package
func doc2() {
	fmt.Println("Faz nada, só testando...")
}

// Doc 3 é a ultima func do package de teste
func Doc3() {
	fmt.Println("Faz nada, só testando...", X)
}