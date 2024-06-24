package main

import "fmt"

func Exercicio8() {

	pessoas := map[string][]string{
		"filipe_rodrigues": {
			"sei lá",
		},
		"ayrton_senna": {
			"pilotar aviao", "andar de carro",
		},
		"nicole_alves": {
			"ver filmes",
		},
	}

	for k, v := range pessoas {
		fmt.Println(k)
		for l, w := range v {
			fmt.Println("\t", l+1, "-", w)
		}
	}

}
