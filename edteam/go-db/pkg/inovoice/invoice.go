package invoice

import (
	"github.com/krlosw9/cursosGo/go-db/pkg/invoiceheader"
	"github.com/krlosw9/cursosGo/go-db/pkg/invoiceitem"
)

type Model struct {
	Header *invoiceheader.Model
	Items  invoiceitem.Models
}

type Storage interface {
	Create(*Model) error
}

// Service of invoice
type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
