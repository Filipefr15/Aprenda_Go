package main

import "fmt"

func main() {
	a := i()

	a()
	a()
	a()
	a()
	fmt.Println(a())
}

func i() func() int {
	x := 0 //basicamente o x está aqui fora, então quando a:=1(), o que o a recebe na verdade é a função de baixo!
	return func() int {
		x++
		return x
	}
}
