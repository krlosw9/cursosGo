package main

import (
	"github.com/krlosw9/cursosGo/go-db-gorm/model"
	"github.com/krlosw9/cursosGo/go-db-gorm/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	// Borrado suave o logico
	// myProduct := model.Product{}
	// myProduct.ID = 2

	// storage.DB().Delete(&myProduct)
	// ----------------------------------------------------
	// products := make([]model.Product, 0)
	// storage.DB().Find(&products)

	// for _, product := range products {
	// 	fmt.Printf("%d - %s\n", product.ID, product.Name)
	// }
	// ----------------------------------------------------
	// Borrado permanente
	myProduct := &model.Product{}
	myProduct.ID = 3
	storage.DB().Unscoped().Delete(myProduct)
}
