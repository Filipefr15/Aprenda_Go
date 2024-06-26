package main

import (
	"fmt"
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
		z := Soma(v.data...)
		if z != v.answer {
			t.Error("Expected:", v.answer, "Got:", z)
		}
	}
}

func ExampleSoma() {
	fmt.Println(Soma(5, 2, -1))
	//Output: 6
}

func TestSoma(t *testing.T) {
	teste := Soma(3, 2, 1)
	resultado := 6
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

func TestMultiplica(t *testing.T) {
	teste := Multiplica(3, 2, 1)
	resultado := 6
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

/*
	func BenchmarkSoma(b *testing.B) {
		tests := []test{
			{data: []int{1, 2, 3}, answer: 6},
			{[]int{10, 11, 12}, 33},
			{[]int{-5, 0, 5, 10}, 10},
		}
		for i := 0; i < b.N; i++ {
			for _, v := range tests {
				Soma(v.data...)
			}
		}
	}
*/
func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Soma(1, 20)
	}
}

func BenchmarkMultiplica(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiplica(10, 10, 10)
	}
}
