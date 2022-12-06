package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	contador int = 0
	wg       sync.WaitGroup
	mu       sync.Mutex
)

func main() {
	totalGoroutines := 1000

	wg.Add(totalGoroutines) // Indicando quantas goroutines serão criadas/'lançadas'
	for i := 0; i < totalGoroutines; i++ {
		go func() {
			mu.Lock()
			lerSalvar := contador // Lendo e salvando o valor atual do 'contador'
			runtime.Gosched()     // Fazendo yield da thread
			lerSalvar++           // Incrementando o valor salvo
			contador = lerSalvar  // Copiando o novo valor incrementado para a variável inicial
			fmt.Println(contador)
			mu.Unlock()

			wg.Done() // Sinalizando o fim da goroutine
		}()
	}

	wg.Wait() // Sinalizando à func main() que deve-se esperar as goroutines serem finalizadas
}
