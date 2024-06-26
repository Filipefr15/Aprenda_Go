package main

import "fmt"

func main() {
	canal := make(chan int)

	go func() {
		canal <- 0
		close(canal)
	}()

	v, ok := <-canal //Aqui v receberá 0, e ok será true, já que 0 é realmente o número passado pelo canal

	fmt.Println(v, ok)

	v, ok = <-canal //Aqui v será 0, mas não sera o mesmo 0 passado pelo canal, então ok sera false

	fmt.Println(v, ok)
}
