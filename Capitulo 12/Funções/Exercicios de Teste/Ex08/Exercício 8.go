package main

import "fmt"

func main() {

	y := retornar()

	y()

}

func retornar() func() {
	return func() {
		fmt.Println("pois é")
	}
}
