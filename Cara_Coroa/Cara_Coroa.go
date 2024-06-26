package main

import (
	"fmt"
	"math/rand"
)

func main() {
	i := rand.Intn(100)
	if i == 0 {
		fmt.Println("MEIO")
	} else if i%2 == 0 {
		fmt.Println("COROA")
	} else {
		fmt.Println("CARA")
	}
}
