package main

import (
	"fmt"
)

//Faça esse código funcionar
/*
func main() {
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}
*/
//Usando uma função anônima auto-executável
func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

//Usando buffer
/*
func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
*/
