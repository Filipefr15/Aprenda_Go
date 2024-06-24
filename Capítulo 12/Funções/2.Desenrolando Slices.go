package main

import "fmt"

func desenslices() {

	si := []int{1, 2, 3, 4, 5}
	somando := somaint(si...)
	fmt.Println(somando)
}

func somaint(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma

}
