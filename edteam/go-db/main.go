package main

import (
	"log"

	invoice "github.com/krlosw9/cursosGo/go-db/pkg/inovoice"
	"github.com/krlosw9/cursosGo/go-db/pkg/invoiceheader"
	"github.com/krlosw9/cursosGo/go-db/pkg/invoiceitem"
	"github.com/krlosw9/cursosGo/go-db/storage"
)

func main() {
	storage.NewPostgresDB()

	storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
	storageInvoice := storage.NewPsqlInvoice(
		storage.Pool(),
		storageHeader,
		storageItems,
	)

	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Cliente error",
		}, Items: invoiceitem.Models{
			&invoiceitem.Model{ProductID: 2},
			&invoiceitem.Model{ProductID: 99},
		},
	}

	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create: %v", err)
	}
}
