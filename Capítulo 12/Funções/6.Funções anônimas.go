package main

import "fmt"

func main() {

//Declarou e chamou a função ao mesmo tempo.
//Descartavel, so vai rodar uma vez
	x := 10

	func(x int) {
		fmt.Println(x, "vezes 10 é: ")
		fmt.Println(x * 10)
	}(x)

}
