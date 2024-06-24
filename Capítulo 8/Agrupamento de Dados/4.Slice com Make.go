package main

import "fmt"

func main() {

	//slice := make([]int, 10, 10)
	slice := new([10]int)[0:10]

	for i := 0; i < cap(slice); i++ {
		slice[i] = 2
	}
	//fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 10)
	slice = append(slice, 10)

	fmt.Println(slice)
	fmt.Println(slice, len(slice), cap(slice))
}
