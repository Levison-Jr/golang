package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	funções := 5
	go manda(20, canal1)
	go outra(funções, canal1, canal2)
	for v := range canal2 {
		fmt.Println(v)
	}
}

func manda(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
	fmt.Println("FUII")
}

func outra(funções int, canal1, canal2 chan int) {
	var wg sync.WaitGroup
	for i := 0; i < funções; i++ {
		fmt.Println("EU:", i)
		wg.Add(1)
		go func() {
			for v := range canal1 {
				canal2 <- trabalho(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(canal2)
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * 2000)
	return n
}
