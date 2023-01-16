package bank

import (
	"log"

	"github.com/ntferr/icash/entities"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

type Service interface {
	NewBank(bank *entities.Bank) error
}

func NewService(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}

func (s service) NewBank(bank *entities.Bank) error {
	tx := s.db.Create(bank)
	if err := tx.Error; err != nil {
		log.Printf("error to create new tuple of bank: %e", err)
	}
	return nil
}
