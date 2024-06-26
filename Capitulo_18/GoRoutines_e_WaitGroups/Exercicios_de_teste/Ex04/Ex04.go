package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador = 0
var mu sync.Mutex

const quantidadeDeGoroutines = 200

func main() {

	criarGoroutines(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total de Goroutines:", quantidadeDeGoroutines,
		"\nTotal do contador:", contador)
}

func criarGoroutines(x int) {
	wg.Add(x)
	for j := 0; j < x; j++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}
}
