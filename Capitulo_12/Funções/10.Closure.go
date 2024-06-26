package main

import "fmt"

//Closure: capturar/cercar um scope para utilizar somente quando quiser

func mainclosure() {

	a := i()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := i()

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func i() func() int {
	x := 0 //basicamente o x está aqui fora, então quando a:=1(), o que o a recebe na verdade é a função de baixo!
	return func() int {
		x++
		return x
	}
}
