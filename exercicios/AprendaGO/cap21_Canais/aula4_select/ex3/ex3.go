package main

import (
	"fmt"
	"time"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	num := []int{1, 2, 3}
	go send(num, par, impar, quit)
	receive(par, impar, quit)
}

func send(num []int, par chan int, impar chan int, quit chan bool) {
	for i := 0; i < len(num); i++ {
		if num[i]%2 == 0 {
			fmt.Println("PAR enviado")
			par <- num[i]
		} else {
			fmt.Println("ÍMPAR enviado")
			impar <- num[i]
		}
	}

	quit <- true
}

func receive(par chan int, impar chan int, quit chan bool) {
	for {
		time.Sleep(3 * time.Second)
		select {
		case value := <-par:
			fmt.Println("Recebido PAR:", value)
		case value := <-impar:
			fmt.Println("Recebido ÍMPAR:", value)
		case value := <-quit: // Se colocar 'case <-quit:' o retorno ocorre sem considerar o 'time.Sleep'
			fmt.Println(value)
			return
		}
	}
}
