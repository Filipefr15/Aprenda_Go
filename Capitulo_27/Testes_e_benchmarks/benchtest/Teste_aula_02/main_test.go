package main

import (
	"testing"
)

//    - struct test, fields: data []int, answer int
//- tests := []test{[]int{}, int}
//- range tests

type test struct {
	data   []int
	answer int
}

func TestSomaEmTabela(t *testing.T) {
	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{[]int{10, 11, 12}, 33},
		{[]int{-5, 0, 5, 10}, 10},
	}
	for _, v := range tests {
		z := soma(v.data...)
		if z != v.answer {
			t.Error("Expected:", v.answer, "Got:", z)
		}
	}
}

func TestSoma(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 6
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

func TestMultiplica(t *testing.T) {
	teste := multiplica(3, 2, 1)
	resultado := 6
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}
