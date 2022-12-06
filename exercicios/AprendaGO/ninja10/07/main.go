package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	go enviar(c)

	for value := range c {
		fmt.Println(value)
	}
}

func enviar(canal chan int) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		fmt.Println("III ==", i)
		wg.Add(1)
		go func(c chan int, in int) {
			for j := 0; j < 5; j++ {
				c <- (j + in)
			}
			wg.Done()
		}(canal, i)
	}
	//fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	close(canal)
}
