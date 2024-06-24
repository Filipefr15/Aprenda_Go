package main

import "fmt"

func main() {
	sabores := []string{"pepperoni", "pantaneira", "dois queijos", "chocolate"}
	//fatia := sabores[1:len(sabores)]
	//fatia := sabores[:] (pega todos)

	//for i := 0; i < len(sabores); i++ {
	//	fmt.Println(sabores[i])
	//}

	sabores = append(sabores[:2], sabores[4:]...) //Isso é uma atribuição de uma nova slice para o mesmo valor, na nova slice o valor da antiga é colocado já sem o item "excluido"

	fmt.Println("Pós delete:")

	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}

}
