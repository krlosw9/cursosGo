package main

import (
	"github.com/krlosw9/cursosGo/go-db-gorm/model"
	"github.com/krlosw9/cursosGo/go-db-gorm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	storage.DB().AutoMigrate(&model.Product{}, &model.InvoiceHeader{}, &model.InvoiceItem{})

}
