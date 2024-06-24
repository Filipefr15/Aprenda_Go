package main

import "fmt"

func mainrecursividade() {
	fmt.Println(fatorial(4))
	fmt.Println(loops(4))
}

func fatorial(num int) int {
	if num == 0 {
		return 1
	}
	if num == 1 {
		return num
	}
	return num * fatorial(num-1)

}

func loops(x int) int {
	if x == 0 {
		return 1
	}
	total := 1
	for x > 1 {
		total *= x
		x--
	}
	return total
}
