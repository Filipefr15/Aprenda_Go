package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func mainex2() {

	pessoaGenerica := pessoa{
		nome:      "Filipe",
		sobrenome: "Rodrigues",
		idade:     24,
	}
	fmt.Println(pessoaGenerica)
	mudeMe(&pessoaGenerica)
	fmt.Println(pessoaGenerica)
}

func mudeMe(p *pessoa) {
	(*p).nome = "Filinpe"
	p.sobrenome = "Rodringues"
	(*p).idade = 42
	//da para usar (*p). ou p. as duas coisas sao iguais.
}
