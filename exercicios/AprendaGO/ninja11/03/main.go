package main

import (
	"fmt"
	"log"
)

type erroEspecial struct {
	soma int
	err error
}

func (e erroEspecial) Error() string {
	return fmt.Sprintf("Atenção! Um erro foi detectado: %v || %v", e.soma, e.err)
}

func chamaErro(e error) {
	log.Fatalln(e)
}

func main() {
	x := 127
	erro := erroEspecial {
		soma: x,
		err: fmt.Errorf("DEU RUIM!"),
	}

	chamaErro(erro)
}
