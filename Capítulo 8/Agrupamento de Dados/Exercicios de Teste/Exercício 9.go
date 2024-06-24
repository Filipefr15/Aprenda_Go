package main

import "fmt"

func main() {

	pessoas := map[string][]string{
		"filipe_rodrigues": []string{
			"sei lá",
		},
		"ayrton_senna": []string{
			"pilotar aviao", "andar de carro",
		},
		"nicole_alves": []string{
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

	for k, v := range pessoas {
		fmt.Println(k)
		for l, w := range v {
			fmt.Println("\t", l+1, "-", w)
		}
	}

}
