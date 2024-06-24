package main

import "fmt"

//assim dá errado! Eles tem números finitos
func main() {
	canal := make(chan int, 1) //mudar o num pra 2 e adicionar um println da certo

	canal <- 42
	canal <- 43

	fmt.Println(<-canal)
	//fmt.Println(<-canal)
}

/*
Assim dá certo, já que passamos o 1 no make (buffer)
func main() {
	canal := make(chan int, 1)

	canal <- 42

	fmt.Println(<-canal)
}

*/
