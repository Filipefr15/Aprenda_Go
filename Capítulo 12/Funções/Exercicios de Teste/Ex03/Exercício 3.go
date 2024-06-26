package main

import "fmt"

func main() {
	defer fmt.Println("Sou o primeiro contador")
	fmt.Println("Sou o segundo contador")
	fmt.Println("Sou o terceiro contador")
	fmt.Println("Sou o quarto contador")
}
