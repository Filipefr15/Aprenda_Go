package main

import "fmt"

func main() {

	si := []int{1, 2, 3, 4, 5}
	somando := soma(si...)
	fmt.Println(somando)
}

func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma

}
