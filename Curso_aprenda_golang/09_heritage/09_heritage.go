package main

import "fmt"

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

type PessoaJuridica struct {
	RazaoSocial string
	cnpj        string
}

func (p PessoaFisica) String() string {
	return fmt.Sprintf("Olá, meu nome é %s %s, eu tenho %d anos, sou %t e meu cpf é: %s.", p.Nome, p.Sobrenome, p.Idade, p.Status, p.cpf)
}

func (p PessoaJuridica) String() string {
	return fmt.Sprintf("Meu CNPJ é: %s e minha razão social é: %s", p.cnpj, p.RazaoSocial)
}

func main() {
	p := PessoaFisica{
		Pessoa{"Filipe", 24, true},
		"Ferreira",
		"12345678901",
	}

	p2 := PessoaJuridica{"1", "2"}
	fmt.Println(p2)
	fmt.Println(p)
}
