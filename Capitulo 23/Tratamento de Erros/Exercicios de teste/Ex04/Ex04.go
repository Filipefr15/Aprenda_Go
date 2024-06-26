package main

import (
	"fmt"
	"log"
	"reflect"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		erroNovo := fmt.Errorf("Deu erro no valor: %v", f)
		fmt.Println(reflect.TypeOf(erroNovo))
		return 0, sqrtError{"lat: sim", "long: sim tbm", erroNovo}
	}
	return 42, nil
}
