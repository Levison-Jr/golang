package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	contador int64 = 0
	wg       sync.WaitGroup
)

func main() {
	totalGoroutines := 1000
	wg.Add(totalGoroutines) // Indicando quantas goroutines serão criadas/'lançadas'

	for i := 0; i < totalGoroutines; i++ {

		go func() {
			atomic.AddInt64(&contador, 1)
			fmt.Println(atomic.LoadInt64(&contador))
			wg.Done() // Sinalizando o fim da goroutine
		}()
	}

	wg.Wait() // Sinalizando à func main() que deve-se esperar as goroutines serem finalizadas
	fmt.Println("Valor Final:", contador)
}
