package main

import "fmt"

func Exercicio6() {
	estados := make([]string, 27, 30)
	estados = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Distrito Federal", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais",
		"Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}
	fmt.Println(len(estados))
	fmt.Println(cap(estados))

	for i := 0; i < len(estados); i++ {
		fmt.Println(estados[i])
	}
}
