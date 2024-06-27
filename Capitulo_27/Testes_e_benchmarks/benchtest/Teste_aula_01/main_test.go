package main

import "testing"

//func soma(i ...int) int

func TestSoma(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 6
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

//Esse teste falha
/*
func TestSoma2(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 7
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}
*/

func TestMultiplica(t *testing.T) {
	teste := multiplica(3, 2, 1)
	resultado := 6
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}
