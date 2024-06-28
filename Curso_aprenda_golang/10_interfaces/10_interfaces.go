package main

import "fmt"

type Document interface {
	Doc() string
}

type Pessoa struct {
	Nome   string
	Idade  int
	Status bool
}

type PessoaFisica struct {
	Pessoa
	Sobrenome string
	cpf       string
}

func (pf PessoaFisica) Doc() string {
	return fmt.Sprintf("Meu cpf é: %s", pf.cpf)
}

func (pf PessoaJuridica) Doc() string {
	return fmt.Sprintf("Meu CNPJ é: %s", pf.cnpj)
}

type PessoaJuridica struct {
	Pessoa
	RazaoSocial string
	cnpj        string
}

func show(d Document) {
	fmt.Println(d)
	fmt.Println(d.Doc())
}

func main() {
	p := PessoaFisica{
		Pessoa{"Filipe", 24, true},
		"Ferreira",
		"12345678901",
	}

	pj := PessoaJuridica{Pessoa{"Filipe", 24, true}, "1", "2"}
	show(p)
	show(pj)
}
