package main

import "fmt"

type figuras2D interface {
	area() float64
}

func calcularArea(f figuras2D) {
	fmt.Println("Area: ", f.area())
}

type cuadrado struct {
	base float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

type rectangulo struct {
	base   float64
	altura float64
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func main() {
	myCuadrado := cuadrado{2}
	myRectangulo := rectangulo{2, 4}

	calcularArea(myCuadrado)
	calcularArea(myRectangulo)

	// Lista de interfaces -> es para tener una lista con datos de diferentes tipos (string, int, float....)
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)

}
