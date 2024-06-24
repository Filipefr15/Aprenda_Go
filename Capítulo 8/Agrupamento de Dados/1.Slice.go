package main

import "fmt"

func main() {
	slice := []string{"banana", "maçã", "jaca"}
	for índice, valor := range slice {
		fmt.Println("No índice", índice, "temos o valor:", valor)
	}

	slice = append(slice, "melão")
	fmt.Println("Fruta extra: ", slice[3])
}
