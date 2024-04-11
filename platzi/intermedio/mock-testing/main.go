package main

import (
	"time"
)

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	id       int
	Position string
}

type FullTimeEmployee struct {
	Person
	Employee
}

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmploye FullTimeEmployee

	e, err := GetEmployeeById(id)
	if err != nil {
		return ftEmploye, err
	}
	ftEmploye.Employee = e

	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmploye, err
	}
	ftEmploye.Person = p

	return ftEmploye, nil
}
