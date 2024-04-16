package main

import (
	"github.com/krlosw9/cursosGo/go-db-gorm/model"
	"github.com/krlosw9/cursosGo/go-db-gorm/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	myProduct := model.Product{}
	myProduct.ID = 3

	storage.DB().Model(&myProduct).Updates(
		model.Product{Name: "Curso de Java", Price: 120},
	)

}
