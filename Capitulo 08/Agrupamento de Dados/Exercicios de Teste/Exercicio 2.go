package main

import "fmt"

func Ex2() {

	slicequalquer := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for _, v := range slicequalquer {
		fmt.Println(v)
	}

	fmt.Printf("%T", slicequalquer)

}
