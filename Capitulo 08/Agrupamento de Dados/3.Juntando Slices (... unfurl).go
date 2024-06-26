package main

import "fmt"

func Jslices() {

	umaslice := []int{1, 2, 3, 4}
	anotherslice := []int{9, 10, 11, 12}

	umaslice = append(umaslice, anotherslice...) //unfurl - desenrola esse negócio e coloca um item de cada vez!

	fmt.Println(umaslice)

}
