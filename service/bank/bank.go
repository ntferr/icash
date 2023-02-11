package bank

import (
	"errors"
	"log"

	"github.com/ntferr/icash/entities"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

type Service interface {
	Get(id string) (*entities.Bank, error)
	Insert(bank *entities.Bank) error
	Update(bank *entities.Bank) error
	Delete(id string) error
}

func NewService(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}

func (s service) Get(id string) (*entities.Bank, error) {
	var err error
	tuple, ok := s.db.Get(id)
	if !ok {
		log.Printf("failed to get bank %s", id)
		err = errors.New("failed to get bank")
		return nil, err
	}

	bank, ok := tuple.(entities.Bank)
	if !ok {
		log.Printf("failed to cast to struct bank")
		err = errors.New("failed to cast to struct bank")
	}

	return &bank, err
}

func (s service) Insert(bank *entities.Bank) error {
	tx := s.db.Create(bank)
	err := tx.Error
	if err != nil {
		log.Printf("failed to create new tuple of bank: %e", err)
	}
	return err
}

func (s service) Update(bank *entities.Bank) error {
	tx := s.db.Update(bank.ID, bank)
	err := tx.Error
	if err != nil {
		log.Printf("failed to update bank %s: %e", bank.ID, err)
	}
	return nil
}

func (s service) Delete(id string) error {
	tx := s.db.Delete(id)
	err := tx.Error
	if err != nil {
		log.Printf("failed to delete bank %s: %e", id, err)
	}
	return err
}
