package main

import "fmt"

func main() {
	var (
		contGre, contInter, contEmp int
		golGre, golInter, repetir   int
	)

	for {
		fmt.Scanln(&golInter, &golGre)

		switch {
		case golGre > golInter:
			contGre++
		case golGre < golInter:
			contInter++
		default:
			contEmp++
		}

		fmt.Println("Novo grenal (1-sim 2-nao)")
		fmt.Scan(&repetir)
		if repetir == 2 {
			break
		}
	}

	fmt.Printf("%d grenais\n", contEmp+contGre+contInter)
	fmt.Printf("Inter:%d\n", contInter)
	fmt.Printf("Gremio:%d\n", contGre)
	fmt.Printf("Empates:%d\n", contEmp)

	switch {
	case contGre > contInter:
		fmt.Println("Gremio venceu mais")
	case contGre < contInter:
		fmt.Println("Inter venceu mais")
	default:
		fmt.Println("Nao houve vencedor")
	}
}
