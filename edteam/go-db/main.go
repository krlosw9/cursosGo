package main

import (
	"log"

	"github.com/krlosw9/cursosGo/go-db/pkg/invoiceheader"
	"github.com/krlosw9/cursosGo/go-db/pkg/invoiceitem"
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

	storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

	if err := serviceInvoiceHeader.Migrate(); err != nil {
		log.Fatalf("invoiceheader.Migrate: %v", err)
	}

	storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)

	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("invoiceitem.Migrate: %v", err)
	}
}
