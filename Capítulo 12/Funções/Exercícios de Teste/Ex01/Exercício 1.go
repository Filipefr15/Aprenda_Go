package main

import "fmt"

func main() {
	fmt.Println(retornarInt())
	fmt.Println(retornarIntString())
}

func retornarInt() int {
	return 2
}

func retornarIntString() (int, string) {
	return 3, "?"
}
