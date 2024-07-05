package main

import "fmt"

func main() {
	hello := make(chan string)

	go func() {
		hello <- "Olá"
	}()

	result := <-hello
	fmt.Println(result)
}
