package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func printString(param1, param2 int, param3 string) {
	fmt.Println(param1, param2, param3)
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func isPalindromo(text string) {
	var reverseText string
	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		reverseText += string(text[i])
	}

	if text == reverseText {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}

}

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

	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	// numeros complejos
	// Complex64 = Real e Imaginario float32
	// Complex128 = Real e Imaginario float64
	// Ejemplo : c:=10 + 8i

	helloMessage := "Hello"
	worldMessage := "world"

	// Println: Salto de Linea Automatico
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	// Con valores seguros
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	// Con valores inseguros
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%v tiene más de %v cursos\n", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos:
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

	printString(1, 2, "parametro3")
	var value1, value2 int = doubleReturn(2)
	var value3, _ int = doubleReturn(6)

	println(value1, value2, value3)

	for i := 0; i < 10; i++ {
		println(i)
	}
	fmt.Println()

	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// counterForEver := 0
	// for {
	// 	fmt.Println(counterForEver)
	// 	counterForEver++
	// }

	// for range
	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range listaNumerosPares {
		fmt.Printf("posicion %d número par: %d \n", i, par)
	}

	// Convertir texto a numero
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value: ", value)

	//SWITCH

	//structure of switch
	auxiliar := 10
	switch auxiliar {
	case 5:
		fmt.Println("Value is 5")
	case 10:
		fmt.Println("Value is 10")
	default:
		fmt.Println("Unknown value")
	}

	//If using a variable is recommended to do this

	switch auxiliar2 := 10; auxiliar2 {
	case 5:
		fmt.Println("Value is 5")
	case 10:
		fmt.Println("Value is 10")
	default:
		fmt.Println("Unknown value")
	}

	//Switch without a condition
	value4 := 50
	switch {
	case value4 < 0:
		fmt.Println("Value is smaller than zero")
	case value > 100:
		fmt.Println("Value is greater than 100")
	default:
		fmt.Println("Value is between 0 and 100")
	}

	fmt.Println() // Salto de linea
	defer fmt.Println("Primero")
	fmt.Println("Segundo")

	for i := 0; i < 10; i++ {
		//continue
		if i == 2 {
			fmt.Println("llegue a 2")
			continue
		}
		//break
		if i == 8 {
			fmt.Println("llegue a 8")
			break
		}
		fmt.Println(i)
	}

	// Array -> son inmutables y se les define cantidad de elementos que va a tener desde el inicio (son mas eficientes)
	fmt.Println() // salto de linea
	var array [4]int
	fmt.Println(array) // [0 0 0 0]
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array)) // [1 2 0 0] 4(longitud del array) 4(capacidad del array)
	// Slice -> son mutables, no requieren que se defina inicialmente la cantidad de elementos que va a tener
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Metodos en slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice) // [0 1 2 3 4 5 6 7]

	// Append nueva lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice) // [0 1 2 3 4 5 6 7 8 9 10]

	fmt.Println() // salto de linea
	isPalindromo("amor A Roma")

	fmt.Println() // salto de linea
	// Maps
	m := make(map[string]int)

	m["jose"] = 21
	m["pepito"] = 14

	fmt.Println(m) // map[jose:21 pepito:14]

	for i, v := range m {
		fmt.Println(i, v)
	}

	value, ok := m["jose"]
	fmt.Println(value, ok)

	value5, ok := m["josep"] // Esto no existe en el Map
	fmt.Println(value5, ok)
}
