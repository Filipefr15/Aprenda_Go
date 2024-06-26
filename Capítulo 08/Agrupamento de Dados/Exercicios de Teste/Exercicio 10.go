package main

import "fmt"

func Exercicio10() {

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

	pessoas["barry_allen"] = []string{"correr", "combater o crime", "fazer piadas"}

	fmt.Println("New slice:")

	delete(pessoas, "ayrton_senna")

	for k, v := range pessoas {
		fmt.Println(k)
		for l, w := range v {
			fmt.Println("\t", l+1, "-", w)
		}
	}

}
