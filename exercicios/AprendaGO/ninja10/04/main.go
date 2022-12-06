package main

import (
	"fmt"
	"time"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q <-chan int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("Aguardando a outra goroutine ler...")
			c <- i
			fmt.Println("Mandei:", i)
		}
		fmt.Println("CABEI...QUERO CAIR FORA!")
		<-q
	}()
	return c
}

func receive(c <-chan int, q chan int) {
	for {
		select {
		case v := <-c:
			time.Sleep(5 * time.Second)
			fmt.Println(v)
		case q <- 0:
			fmt.Println("ENTÃO VUMBORA!")
			return
		}
	}
}
