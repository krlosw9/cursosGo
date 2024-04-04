package main

import "fmt"

type Employee struct {
	id   int
	name string
}

// visible function
func (e *Employee) setId(id int) {
	e.id = id
}

func (e *Employee) setName(name string) {
	e.name = name
}

func (e *Employee) getId() int {
	return e.id
}

func (e *Employee) getName() string {
	return e.name
}

func main() {
	e := Employee{}
	fmt.Println(e)
	e.id = 1
	e.name = "Name"
	e.setId(2)
	e.setName("Name2")
	fmt.Println(e)
	fmt.Printf("id: %d, Name: %s\n", e.getId(), e.getName())
}
