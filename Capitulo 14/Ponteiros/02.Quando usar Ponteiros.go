package main

import "fmt"

func main() {

	x := 0

	//estarecebeovalor(x)
	estarecebeumponteiro(&x)
	fmt.Println(x)

}

func estarecebeovalor(x int) {
	x++
	fmt.Println("Na função:", x)
}

func estarecebeumponteiro(x *int) {
	*x++
	fmt.Println("Na função:", *x)
}
