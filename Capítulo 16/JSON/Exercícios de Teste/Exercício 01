//
//
//LEMBRAR DA LETRA MAIÚSCULA NOS NOMES, SÓ EXPORTA PARA JSON AS MAIUSCULAS!!!!
//
//
package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	b, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
	fmt.Println(users)
}
