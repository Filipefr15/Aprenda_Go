package main

import "fmt"

func main() {
	canalPrincipal := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			canalPrincipal <- i
		}
		close(canalPrincipal)
	}()

	for v := range canalPrincipal {
		fmt.Println("Número retirado do canal:", v)
	}

}

/*
Solução da professora
func main() {
	c := make(chan int)
	go botalá(c)
	for v := range cc {
		fmt.Println(v)
	}

func botalá(c chan int){
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

}
*/
