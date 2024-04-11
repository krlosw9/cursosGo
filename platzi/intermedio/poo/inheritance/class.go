package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

// func GetMessage(p Person) {
// 	fmt.Printf("%s with age %d\n", p.name, p.age)
// }

func main() {
	ftEmploye := FullTimeEmployee{}
	ftEmploye.id = 1
	ftEmploye.name = "Name1"
	ftEmploye.age = 20
	fmt.Println(ftEmploye)
	// GetMessage(ftEmploye) -> no se puede el polimorfismo en Go porque go tiene composicion no herencia

}
