package main

import (
	"fmt"
	"math/rand"
	"sort"
)

const (
	QtdeNumeros = 60
)

func main() {
	fmt.Println("Aqui estão seus números:")
	for i := 0; i < 6; i++ {
		repeat()
	}
}

func repeat() {
	slice := []int{rand.Intn(QtdeNumeros), rand.Intn(QtdeNumeros), rand.Intn(QtdeNumeros), rand.Intn(QtdeNumeros), rand.Intn(QtdeNumeros), rand.Intn(QtdeNumeros)}
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if slice[i] == slice[j] && i != j {
				slice[i] = rand.Intn(QtdeNumeros)
				i, j = 0, 0
			}
			if slice[i] == 0 {
				slice[i] = rand.Intn(QtdeNumeros)
				i, j = 0, 0
			}
		}
	}
	sort.Ints(slice)
	fmt.Println(slice)
}
