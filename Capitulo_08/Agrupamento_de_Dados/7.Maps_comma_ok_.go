package main

import "fmt"

func maps() {

	amigos := map[string]int{
		"alfredo": 12345644,
		"filipe":  9995487,
	}

	//fmt.Println(amigos["filipe"])

	amigos["andrei"] = 959595

	fmt.Println(amigos)

	teste, ok := amigos["anddrei"]

	fmt.Println(teste, ok)
}
