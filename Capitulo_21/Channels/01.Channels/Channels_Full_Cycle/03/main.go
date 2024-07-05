package main

import (
	"fmt"
	"time"
)

func main() {
	hello := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		hello <- "Olá"
	}()

	select {
	case x := <-hello:
		fmt.Println(x)
	default:
		fmt.Println("Padrão")
	}

	time.Sleep(time.Second * 5)

}
