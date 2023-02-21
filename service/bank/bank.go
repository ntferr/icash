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
	GetAll() ([]*entities.Bank, error)
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

func (s service) GetAll() ([]*entities.Bank, error) {
	var banks []*entities.Bank
	result := s.db.Find(&banks)
	if err := result.Error; err != nil {
		return nil, err
	}

	return banks, nil
}

func (s service) Get(id string) (*entities.Bank, error) {
	var bank *entities.Bank

	tx := s.db.Find(&bank, "id = ?", id)
	if err := tx.Error; err != nil {
		log.Printf("failed to get bank %s: %e", id, err)
		return nil, err
	}

	return bank, nil
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
	tx := s.db.Where("id = ?", bank.ID).Updates(&bank)
	err := tx.Error
	if err != nil {
		log.Printf("failed to update bank %s: %e", bank.ID, err)
	}

	return nil
}

func (s service) Delete(id string) error {
	var bank entities.Bank
	tx := s.db.Where("id = ?", id).Delete(&bank)
	err := tx.Error
	if err != nil {
		log.Printf("failed to delete bank %s: %e", id, err)
	}

	return err
}
