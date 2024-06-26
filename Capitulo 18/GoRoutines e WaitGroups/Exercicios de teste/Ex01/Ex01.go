package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		fmt.Println("Sou a primeira goroutine")
		wg.Done()
	}()
	go func() {
		fmt.Println("Sou a segunda gorountine")
		wg.Done()
	}()
	wg.Wait()
}
