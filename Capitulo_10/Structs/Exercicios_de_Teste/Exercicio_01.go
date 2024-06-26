package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func exercicio01() {

	pessoa1 := pessoa{
		nome:      "Filipe",
		sobrenome: "Rodrigues",
		sabores:   []string{"Ninho trufado", "Cream cheese", "Morango"},
	}

	pessoa2 := pessoa{
		nome:      "Filipe2",
		sobrenome: "Rodrigues2",
		sabores:   []string{"Pizza", "Pêssego", "Abobrinha"},
	}

	for i := 0; i < len(pessoa1.sabores); i++ {
		fmt.Println(pessoa1.nome, pessoa1.sobrenome, " sabor ", i+1, "é: ", pessoa1.sabores[i])
	}
	for i := 0; i < len(pessoa2.sabores); i++ {
		fmt.Println(pessoa2.nome, pessoa2.sobrenome, " sabor ", i+1, "é: ", pessoa2.sabores[i])
	}

	//Utilizando range!
	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.sabores {
		fmt.Println("\t", v)
	}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, v := range pessoa2.sabores {
		fmt.Println("\t", v)
	}

}
