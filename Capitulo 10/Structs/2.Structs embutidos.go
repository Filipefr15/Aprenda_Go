package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func mainsStrucsEmbutidos() {
	pessoa1 := pessoa{
		nome:      "Filipe",
		sobrenome: "Rodrigues",
	}

	pessoa2 := pessoa{
		nome:      "Marcelo",
		sobrenome: "Não informado",
	}

	profissional1 := profissional{
		pessoa: pessoa{ //Também funciona se colocar apenas "pessoa1"
			nome:      "Filipe",
			sobrenome: "Rodrigues",
		},
		titulo:  "Dev Trainee",
		salario: 0,
	}

	fmt.Println(pessoa1, pessoa2, profissional1)
}
