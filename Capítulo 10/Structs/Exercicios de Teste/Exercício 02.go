package main

import "fmt"

type pessoas struct {
	nome      string
	sobrenome string
	sabores   []string
}

func exercicio02() {

	meumapa := make(map[string]pessoas)

	//meumapa2 := map[string]pessoas{
	//	"Rodrigues": pessoa{
	//		nome:      "Filipe",
	//		sobrenome: "Rodrigues",
	//		sabores:   []string{"Ninho trufado", "Cream cheese", "Morango"},
	//	},
	//	"Rodrigues2": pessoas{
	//		nome:      "Filipe2",
	//		sobrenome: "Rodrigues2",
	//		sabores:   []string{"Pizza", "Pêssego", "Abobrinha"},
	//	},
	//}

	meumapa["Rodrigues"] = pessoas{
		nome:      "Filipe",
		sobrenome: "Rodrigues",
		sabores:   []string{"Ninho trufado", "Cream cheese", "Morango"},
	}

	meumapa["Rodrigues2"] = pessoas{
		nome:      "Filipe2",
		sobrenome: "Rodrigues2",
		sabores:   []string{"Pizza", "Pêssego", "Abobrinha"},
	}

	for _, valor := range meumapa {
		fmt.Println("Me chamo ", valor.nome, "e meus sabores de sorvete favoritos são: ")
		for _, valor := range valor.sabores {
			fmt.Println("\t", valor)
		}
	}

	//for _, valor := range meumapa2 {
	//	fmt.Println(valor)
	//}

}
