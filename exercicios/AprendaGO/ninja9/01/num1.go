package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	totalGoroutines := 10
	fmt.Println(runtime.NumGoroutine())

	wg.Add(2*totalGoroutines)
	for i := 0; i < totalGoroutines; i++ {
		go func1()
	}
	for i := 0; i < totalGoroutines; i++ {
		go func2()
	}

	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}

func func1() {
	fmt.Println("Olá, FuncUM!")
	wg.Done()
}

func func2() {
	fmt.Println("Olá, FuncDOIS!")
	wg.Done()
}