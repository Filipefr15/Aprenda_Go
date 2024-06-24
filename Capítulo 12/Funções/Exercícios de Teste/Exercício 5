package main

import "fmt"

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

func (r circulo) area() {
	resultado := 3.14 * (r.raio * r.raio)
	fmt.Println("Área do círculo:", resultado)
}

func (l quadrado) area() {
	resultado := l.lado * l.lado
	fmt.Println("Área do quadrado:", resultado)
}

type figura interface {
	area()
}

func info(x figura) {
	x.area()
}

func main() {

	quadrado1 := quadrado{
		lado: 10.2,
	}

	circulo1 := circulo{
		raio: 10.5,
	}

	info(quadrado1)
	info(circulo1)

}
