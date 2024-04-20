package storage

import (
	"fmt"

	"github.com/krlosw9/cursosGo/api-go/class-3/model"
)

// Memory
type Memory struct {
	currentID int
	Persons   map[int]model.Person
}

func NewMemory() Memory {
	persons := make(map[int]model.Person)

	return Memory{
		currentID: 0,
		Persons:   persons,
	}
}

func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	m.currentID++
	m.Persons[m.currentID] = *person

	return nil
}

func (m *Memory) Update(ID int, person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExist)
	}

	m.Persons[ID] = *person

	return nil
}

func (m *Memory) Delete(ID int) error {
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExist)
	}

	delete(m.Persons, ID)

	return nil
}
