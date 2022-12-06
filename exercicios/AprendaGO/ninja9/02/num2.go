package main

import "fmt"

type humanos interface {
	falar()
}

type pessoa struct {
	nome string
	idade int
	cargo string
}

func (p *pessoa) falar() {
	fmt.Printf("Olá, eu sou o %v e sou %v na empresa. Como posso ajudá-lo(a)?", (*p).nome, (*p).cargo)
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {
	p1 := pessoa {
		nome: "José",
		idade: 34,
		cargo: "Gerente",
	}

	dizerAlgumaCoisa(&p1)
}