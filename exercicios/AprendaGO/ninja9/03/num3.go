package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	contador int = 0
	wg       sync.WaitGroup
)

func main() {
	totalGoroutines := 1000

	wg.Add(totalGoroutines) // Indicando quantas goroutines serão criadas/'lançadas'
	for i := 0; i < totalGoroutines; i++ {
		go func() {
			lerSalvar := contador // Lendo e salvando o valor atual do 'contador'
			runtime.Gosched()     // Fazendo yield da thread
			lerSalvar++           // Incrementando o valor salvo
			contador = lerSalvar  // Copiando o novo valor incrementado para a variável inicial
			fmt.Println(contador)
			wg.Done() // Sinalizando o fim da goroutine
		}()
	}

	wg.Wait() // Sinalizando à func main() que deve-se esperar as goroutines serem finalizadas
}

// - Utilizando goroutines, crie um programa incrementador:
// - Tenha uma variável com o valor da contagem
// - Crie um monte de goroutines, onde cada uma deve:
// 	- Ler o valor do contador
// 	- Salvar este valor
// 	- Fazer yield da thread com runtime.Gosched()
// 	- Incrementar o valor salvo
// 	- Copiar o novo valor para a variável inicial
// - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
// - Demonstre que há uma condição de corrida utilizando a flag -race
