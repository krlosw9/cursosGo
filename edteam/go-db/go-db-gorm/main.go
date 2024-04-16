package main

import (
	"fmt"

	"github.com/krlosw9/cursosGo/go-db-gorm/model"
	"github.com/krlosw9/cursosGo/go-db-gorm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	myProduct := model.Product{}
	storage.DB().First(&myProduct, 3)
	fmt.Println(myProduct)

}
