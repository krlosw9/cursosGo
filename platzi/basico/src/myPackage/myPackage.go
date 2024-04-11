package mypackage

import "fmt"

// struct publico
type CarPublic struct {
	Brand string // esto es un atributo publico
	Year  int
}

// struct privado == class private
type carPrivate struct {
	brand string // esto es un atributo privado
	year  int
}

// Esta es una funcion publica por iniciar en mayuscula
func PrintMethod() {
	var myCar carPrivate
	myCar.brand = "Carro privado"
	myCar.year = 2021
	fmt.Println(myCar)
	privateFunction()
}

func privateFunction() {
	fmt.Println("Esta es una funcion privada y solo va ser llamada desde myPackage")
}
