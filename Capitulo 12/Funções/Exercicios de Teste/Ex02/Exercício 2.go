package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 55, 25}
	fmt.Println(retornarInt(slice...))
	fmt.Println(retornarSliceInt(slice))
}

func retornarInt(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func retornarSliceInt(x []int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}
