package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	wg.Add(2)

	go meuLoop(10, c)
	go prints(c)

	wg.Wait()
}

func meuLoop(total int, send chan<- int) {
	for i := 0; i < total; i++ {
		send <- i
	}
	close(send)
	wg.Done()
}

func prints(r <-chan int) {
	for value := range r {
		fmt.Println("Valor recebido:", value)
	}
	wg.Done()
}
