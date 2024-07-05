package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	~int | int64 | float64
} //o ~ faz com que o go aceite o MyNumber

type MyNumber int

func SomaGenerica[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// comparable só compara igualdade!
func Soma[T comparable](number1 T, number2 T) T {
	if number1 == number2 {
		return number1
	}
	return number2
}

func Max[T constraints.Ordered](number1 T, number2 T) T {
	if number1 > number2 {
		return number1
	}
	return number2
}

type stringer interface {
	String() string
}

func concat[T stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		result += val.String()
	}
	return result
}

type MyString string

func (m MyString) String() string {
	return string(m)
}

func main() {

	fmt.Println(concat([]MyString{"a", "b", "c"}))

	var x, y, z MyNumber
	x, y, z = 1, 2, 3

	fmt.Println(SomaGenerica(map[string]MyNumber{"i": x, "c": y, "d": z}))
	fmt.Println(SomaGenerica(map[string]int64{"i": 10, "c": 20, "d": 15}))
	fmt.Println(SomaGenerica(map[string]float64{"i": 10.2, "c": 20.5, "d": 15}))
}
