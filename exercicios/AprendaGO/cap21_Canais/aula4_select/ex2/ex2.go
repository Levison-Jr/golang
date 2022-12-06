package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan int)
	quit := make(chan int)

	go recebeQuit(canal, quit)
	enviaPraCanal(canal, quit)
	//time.Sleep(1 * time.Second) // Sem esse 'time' o 'QUITEI' não tem tempo de ser 'printado'...
}

func recebeQuit(canal chan int, quit chan int) {
	for i := 0; i < 2; i++ {
		fmt.Println("Lendo...(Aguardando dados...)")
		fmt.Println("Recebido:", <-canal)
		fmt.Println("...")
	}

	quit <- 0
	fmt.Println("QUITEI...")
}

func enviaPraCanal(canal chan int, quit chan int) {
	qualquercoisa := 1
	for {
		time.Sleep(1 * time.Second)
		select {
		case canal <- qualquercoisa:
			qualquercoisa++
		case <-quit:
			return
		}
	}
}
