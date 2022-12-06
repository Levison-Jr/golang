package main

import (
	"fmt"
	"github.com/Levison-Jr/hello-world"
)

func main() {
	var (
		n1, n2, n3, n4 float64
	)
	hello.PrintHello()
	fmt.Scanln(&n1, &n2, &n3, &n4)
	media := (2*n1 + 3*n2 + 4*n3 + 1*n4) / (2 + 3 + 4 + 1)
	fmt.Printf("Media: %.1f\n", media)

	if media >= 7.0 { //Média maior que 7.0 == Aprovado
		fmt.Printf("Aluno aprovado.\n")
	} else if media < 5.0 { //Média menor que 5.0 == Reprovado
		fmt.Printf("Aluno reprovado.\n")
	} else if media >= 5.0 && media <= 6.9 { //Média entre 5.0 e 6.9 [5.0,6.9] == Exame
		fmt.Printf("Aluno em exame.\n")

		notaExame := 0.0
		fmt.Scanln(&notaExame)
		fmt.Printf("Nota do exame: %.1f\n", notaExame)

		mediaFinal := (media + notaExame) / 2
		if mediaFinal >= 5.0 {
			fmt.Printf("Aluno aprovado.\n")
		} else {
			fmt.Printf("Aluno reprovado.\n")
		}

		fmt.Printf("Media final: %.1f\n", mediaFinal)
	}

}
