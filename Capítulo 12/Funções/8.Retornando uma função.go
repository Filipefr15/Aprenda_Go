package main

import "fmt"

func main() {

	//Função como expressão!!

	x := retornaumafuncao()
	y := x(3)
	fmt.Println(y)
	fmt.Println(retornaumafuncao()(3))

}

func retornaumafuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
