package main

import "fmt"

type endpoint struct {
	Caminho string `json:"caminho"`
	Metodo  string `json:"metodo"`
	Backend string `json:"backend"`
}

var endPointsMap map[string]map[string]endpoint

func main() {
	// Inicializando o mapa
	endPointsMap = make(map[string]map[string]endpoint)

	// Inserindo informações no mapa
	if endPointsMap["/caminho1"] == nil {
		endPointsMap["/caminho1"] = make(map[string]endpoint)
	}
	endPointsMap["/caminho1"]["GET"] = endpoint{
		Caminho: "/caminho1",
		Metodo:  "GET",
		Backend: "backend1",
	}
	endPointsMap["/caminho1"]["POST"] = endpoint{
		Caminho: "/caminho1",
		Metodo:  "POST",
		Backend: "backend1",
	}

	// Exibindo o conteúdo do mapa
	for caminho, metodos := range endPointsMap {
		for metodo, value := range metodos {
			fmt.Printf("Caminho: %s, Metodo: %s, Value: %+v\n", caminho, metodo, value)
		}
	}
}
