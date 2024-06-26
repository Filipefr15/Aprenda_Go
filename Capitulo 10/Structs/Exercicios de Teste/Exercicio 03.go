package main

import "fmt"

type veiculo struct {
	portas string
	cor    string
}

type caminhonete struct {
	veiculo
	tracao4rodas bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func exercicio03() {
	toro := caminhonete{
		veiculo: veiculo{
			portas: "quatro",
			cor:    "vermelha",
		},
		tracao4rodas: true,
	}

	hb20 := sedan{
		veiculo: veiculo{
			portas: "quatro",
			cor:    "prata",
		},
		modeloLuxo: false,
	}

	fmt.Println(toro)
	fmt.Println(hb20)
	fmt.Println(toro.cor)
	fmt.Println(hb20.modeloLuxo)

}
