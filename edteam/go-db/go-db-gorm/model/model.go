package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string  `gorm:"type:varchar(100); not null"`
	Observation  *string `gorm:"type:varchar(100)"`
	Price        int     `gorm:"not null"`
	InvoiceItems []InvoiceItem
}

type InvoiceHeader struct {
	gorm.Model
	Client       string `gorm:"type:varchar(100); not null"`
	InvoiceItems []InvoiceItem
}

type InvoiceItem struct {
	gorm.Model
	InvoiceHeaderID uint
	ProductID       uint
}
