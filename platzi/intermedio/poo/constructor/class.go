package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{id: id, name: name, vacation: vacation}
}

func main() {
	// 1. Crear una struct con zero value
	e := Employee{}
	fmt.Println(e)

	// 2. Asignarle valores a las propiedades.
	e2 := Employee{id: 1, name: "Name", vacation: false}
	fmt.Println(e2)

	// 3. Usando la palabra reservada new. (Retorna una referencia)
	e3 := new(Employee)
	e3.id = 2
	e3.name = "Name2"
	fmt.Println(e3)
	fmt.Println(*e3)

	// 4.
	e4 := NewEmployee(3, "Name3", true)
	fmt.Println(e4)
	fmt.Println(*e4)
}
