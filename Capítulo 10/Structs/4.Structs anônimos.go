package main

import "fmt"

func mainStrucsAnonimos() {

	x := struct {
		nome  string
		idade int
	}{
		nome:  "Mário",
		idade: 50,
	}
	fmt.Println(x)
}
