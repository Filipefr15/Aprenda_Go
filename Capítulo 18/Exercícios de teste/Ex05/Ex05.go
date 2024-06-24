package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var contador int32 = 0

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
			atomic.AddInt32(&contador, 1)
			v := atomic.LoadInt32(&contador)
			runtime.Gosched()
			fmt.Println(v)
			wg.Done()
		}()
	}
}
