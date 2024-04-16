package main

import (
	"github.com/krlosw9/cursosGo/go-db-gorm/model"
	"github.com/krlosw9/cursosGo/go-db-gorm/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	invoice := model.InvoiceHeader{
		Client: "Alvaro Felipe",
		InvoiceItems: []model.InvoiceItem{
			{ProductID: 1},
			{ProductID: 2},
		},
	}

	storage.DB().Create(&invoice)
}
