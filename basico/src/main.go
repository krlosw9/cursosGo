package main

import "fmt"

func main() {
	fmt.Println("Hola mundo")

	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1416

	fmt.Println("Pi: ", pi)
	fmt.Println("Pi2: ", pi2)

	//Declaracion de variables enteras
	base := 12 // Crea la variable y el runtime elige el tipo de dato
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//  Zero values
	var a int     // 0
	var b float64 // 0
	var c string  // ''
	var d bool    // false

	fmt.Println(a, b, c, d)

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)
}
