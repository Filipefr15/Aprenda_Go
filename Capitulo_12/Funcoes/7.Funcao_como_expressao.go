package main

import "fmt"

func mainexpressions() {

	//Função como expressão!!

	x := 10

	y := func(x int) {
		fmt.Println(x, "vezes 10 é: ")
		fmt.Println(x * 10)
	}

	y(x)

}
