package main

import (
	"github.com/krlosw9/cursosGo/go-db-gorm/model"
	"github.com/krlosw9/cursosGo/go-db-gorm/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	product1 := model.Product{
		Name:  "Curso de go",
		Price: 120,
	}

	obs := "Testing with go"
	product2 := model.Product{
		Name:        "Curso de testing",
		Price:       150,
		Observation: &obs,
	}
	product3 := model.Product{
		Name:  "Curso de python",
		Price: 200,
	}
	storage.DB().Create(&product1)
	storage.DB().Create(&product2)
	storage.DB().Create(&product3)
}
