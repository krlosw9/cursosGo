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

	m := &product.Model{
		ID:          90,
		Name:        "Curso de go",
		Observation: "Observacion del curso de go",
		Price:       50,
	}

	err := serviceProduct.Update(m)
	if err != nil {
		log.Fatalf("product.Update: %v", err)
	}
}
