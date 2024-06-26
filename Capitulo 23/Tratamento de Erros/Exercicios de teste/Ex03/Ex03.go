//- Crie um struct "erroEspecial" que implemente a interface builtin.error.
//- Crie uma função que tenha um valor do tipo error como parâmetro.
//- Crie um valor do tipo "erroEspecial" e passe-o para a função da instrução anterior.

package main

import "fmt"

type erroEspecial struct {
	erro string
}

func (erro erroEspecial) Error() string {
	return "Não deveria colocar o nome erro tantas vezes!"
}

func main() {
	erro := erroEspecial{"Inacreditável!"}
	tipoErroParam(erro)
}

func tipoErroParam(x error) {
	fmt.Println(x.(erroEspecial).erro, x)
}
