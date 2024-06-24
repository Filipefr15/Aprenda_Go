package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

//criando o método falar para este tipo (pessoa) que tenha um receiver ponteiro (*pessoa)
func (p *pessoa) falar() {
	fmt.Println(p.nome, "consegue falar")
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {

	pessoa1 := pessoa{"Filipe", "Ferreira", 24}
	(&pessoa1).falar()         //Funciona
	pessoa1.falar()            //Esse é um atalho de (&pessoa1)
	dizerAlgumaCoisa(&pessoa1) //Funciona
	//dizerAlgumaCoisa(pessoa1)  //Não funciona!!
}
