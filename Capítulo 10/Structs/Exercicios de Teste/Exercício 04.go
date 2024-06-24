package main

import "fmt"

func main() {

	x := struct {
		nome      string
		vontade   int
		ondecomer []string
		vaibemcom map[string]string
	}{
		nome:      "Pizza",
		vontade:   8,
		ondecomer: []string{"Pizzaria", "Restaurantes", "Itália"},
		vaibemcom: map[string]string{
			"de manhã": "com suco",
			"de tarde": "com refrigerante",
			"de noite": "com vinho",
		},
	}
	fmt.Println(x)

}
