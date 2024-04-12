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

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
}
