package main

import "fmt"

func main() {
	canal := make(chan int)
	go enviar(canal)
	receber(canal)
}

func enviar(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func receber(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
