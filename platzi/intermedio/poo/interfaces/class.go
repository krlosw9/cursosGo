package main

import "fmt"

// Interface
type PrintInfo interface {
	getMessage() string
}

// Expone un metodo que utiliza los metodos de quienes implementan la interface
func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

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
	endDate string
}

// Aqui esta implementando la interface para FullTimeEmployee
func (ft FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

type TemporaryEmploye struct {
	Person
	Employee
	taxRate int
}

func (te TemporaryEmploye) getMessage() string {
	return "Temporary Employee"
}

func main() {
	ftEmploye := FullTimeEmployee{}
	ftEmploye.name = "FullName"
	ftEmploye.age = 20
	ftEmploye.id = 1
	ftEmploye.endDate = "20/10/2020"
	fmt.Println(ftEmploye)
	getMessage(ftEmploye)

	tEmployee := TemporaryEmploye{}
	tEmployee.name = "Temporary Name"
	tEmployee.age = 21
	tEmployee.id = 2
	tEmployee.taxRate = 11
	fmt.Println(tEmployee)
	getMessage(tEmployee)

}
