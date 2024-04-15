package invoiceitem

import (
	"database/sql"
	"time"
)

// Model of invoiceitem
type Model struct {
	ID              uint
	InvoiceHeaderID uint
	productID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Models slice of Models
type Models []*Model

type Storage interface {
	Migrate() error
	CreateTx(*sql.Tx, uint, Models) error
}

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
