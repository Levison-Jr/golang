package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var (
	cont []int = make([]int, 10)
	wg sync.WaitGroup
	mu sync.Mutex
)

func main() {
	var (
		A, B int
	)

	for {
		fmt.Scanln(&A, &B)
		if A == 0 && B == 0 {
			break
		}
		for i := A; i <= B; i++ {
			s := strconv.FormatInt(int64(i), 10)
			wg.Add(1)
			go contar(s)
			if runtime.NumGoroutine() == 50 {
				wg.Wait()
			}
		}

		wg.Wait()
		fmt.Println(cont)
		cont = make([]int, 10)
	}
}

func contar(s string) {
	sb := []byte(s)
	for _, value := range sb {
		switch value {
		case 48:
			mu.Lock()
			cont[0]++
			mu.Unlock()
		case 49:
			mu.Lock()
			cont[1]++
			mu.Unlock()
		case 50:
			mu.Lock()
			cont[2]++
			mu.Unlock()
		case 51:
			mu.Lock()
			cont[3]++
			mu.Unlock()
		case 52:
			mu.Lock()
			cont[4]++
			mu.Unlock()
		case 53:
			mu.Lock()
			cont[5]++
			mu.Unlock()
		case 54:
			mu.Lock()
			cont[6]++
			mu.Unlock()
		case 55:
			mu.Lock()
			cont[7]++
			mu.Unlock()
		case 56:
			mu.Lock()
			cont[8]++
			mu.Unlock()
		case 57:
			mu.Lock()
			cont[9]++
			mu.Unlock()
		}
	}
	wg.Done()
}
