package main

import "fmt"

func Exercicio7() {
	multislices := [][]string{
		{"Filipe", "Rodrigues", "Filmes"},
		{"Filipe2", "Rodrigues2", "Tênis"},
		{"Filipe3", "Rodrigues3", "Basquete"},
	}

	for _, v := range multislices {
		fmt.Println(v)
	}

}
