package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	canalPrincipal := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				canalPrincipal <- i
			}
		}()
	}

	for i := 0; i < 100; i++ {
		fmt.Println("Num retirado:", <-canalPrincipal, "I:", i)
	}

}
