package main

import (
	"fmt"
	"time"
)

func main() {
	fila := make(chan int)

	go func() {
		i := 0
		for {
			time.Sleep(time.Second)
			fila <- i
			i++
		}
	}()

	for x := range fila {
		fmt.Println(x)
	}
}
