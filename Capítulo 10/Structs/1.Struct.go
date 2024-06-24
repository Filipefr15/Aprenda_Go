package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func mainStruct() {
	cliente1 := cliente{
		nome:      "Filipe",
		sobrenome: "Rodrigues",
		fumante:   false,
	}

	cliente2 := cliente{
		nome:      "Marcelo",
		sobrenome: "Não informado",
		fumante:   true,
	}

	cliente3 := cliente{"Rubens", "Cervejaria", true}

	fmt.Println(cliente1, cliente2, cliente3)
}
