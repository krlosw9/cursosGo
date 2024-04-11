package main

import (
	"reflect"
	"testing"
)

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						id:       1,
						Position: "CEO",
					}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						DNI:  "1",
						Name: "John Doe",
						Age:  35,
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Age:  35,
					DNI:  "1",
					Name: "John Doe",
				},
				Employee: Employee{
					id:       1,
					Position: "CEO",
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeById(test.id, test.dni)

		if err != nil {
			t.Errorf("Error when getting Employee")
		}

		// if ft.Age != test.expectedEmployee.Age {
		// 	t.Errorf("GetFullTimeEmployeeById was incorrect, got %d expected %d", ft.Age, test.expectedEmployee.Age)
		// }
		if !reflect.DeepEqual(ft, test.expectedEmployee) {
			t.Errorf("Not equal employees, \nExpected: %v\nGot: %v", ft, test.expectedEmployee)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
