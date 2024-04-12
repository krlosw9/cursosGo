package invoiceitem

import "time"

// Model of invoiceitem
type Model struct {
	ID              uint
	InvoiceHeaderID uint
	productID       uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Storage interface {
	Migrate() error
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
