package main

import "fmt"

func main() {

	//defer quer dizer "deixa pra depois", "deixa pra ultima hora"
	//defer é "first in last out"

	//exemplo simples:

	defer fmt.Println("Vim primeiro")
	fmt.Println("Vim depois")
}
