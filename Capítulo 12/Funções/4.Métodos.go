package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "diz bom dia!")
}

func mainmetodos() {

	//Método é uma função anexada a um tipo
	//func (receiver) identifier(parameters) (returns) { code }
	//o método tem a ver com a função "receiver" (receiver diz o contexto da função)
	//Pensar no receiver como referência ao tipo criado (dica de um viewer no yt)

	filipe := pessoa{"Filipe", 24}
	filipe.oibomdia()

}
