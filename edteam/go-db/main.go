package main

import (
	"log"

	"github.com/krlosw9/cursosGo/go-db/pkg/product"
	"github.com/krlosw9/cursosGo/go-db/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	err := serviceProduct.Delete(99)
	if err != nil {
		log.Fatalf("product.Delete: %v", err)
	}
}
