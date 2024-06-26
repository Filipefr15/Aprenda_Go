package main

import "fmt"

func main() {

	x := 10

	func(x int) {
		fmt.Println(x, "vezes 10 é: ")
		fmt.Println(x * 10)
	}(x)

}
