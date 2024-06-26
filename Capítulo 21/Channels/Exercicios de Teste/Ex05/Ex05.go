package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	v, ok := <-c
	if !ok {
		fmt.Println("PROBLEMA!!!!!!", v)
	} else {
		fmt.Println("Tudo ok, olha ai:", v)
	}

	close(c)

	v, ok = <-c
	if !ok {
		fmt.Println("PROBLEMA!!!!!!", v)
	} else {
		fmt.Println("Tudo ok, olha ai:", v)
	}
}
