package main

import "fmt"

func main() {
	canal := make(chan int)

	go func() {
		canal <- 42 //não funcionaria se não estivesse dentro da função
		//é como uma "passagem" de bastão, só funciona se estiver passando de um
		//"canal" para o outro! (nesse caso, de dentro da func para fora)
	}()

	fmt.Println(<-canal)
}
