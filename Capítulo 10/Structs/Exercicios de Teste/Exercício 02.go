package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	meumapa := make(map[string]pessoa)

	//meumapa2 := map[string]pessoa{
	//	"Rodrigues": pessoa{
	//		nome:      "Filipe",
	//		sobrenome: "Rodrigues",
	//		sabores:   []string{"Ninho trufado", "Cream cheese", "Morango"},
	//	},
	//	"Rodrigues2": pessoa{
	//		nome:      "Filipe2",
	//		sobrenome: "Rodrigues2",
	//		sabores:   []string{"Pizza", "Pêssego", "Abobrinha"},
	//	},
	//}

	meumapa["Rodrigues"] = pessoa{
		nome:      "Filipe",
		sobrenome: "Rodrigues",
		sabores:   []string{"Ninho trufado", "Cream cheese", "Morango"},
	}

	meumapa["Rodrigues2"] = pessoa{
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
