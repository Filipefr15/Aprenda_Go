package main

import "fmt"

func Ex1() {

	arrayqualquer := [5]int{10, 20, 30, 40, 50}

	for _, v := range arrayqualquer {
		fmt.Println(v)
	}

	fmt.Printf("%T", arrayqualquer)

}
