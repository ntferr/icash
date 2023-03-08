package bank

import (
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
	tx := s.db.Find(&banks)
	if err := tx.Error; err != nil {
		return nil, err
	}

	return banks, nil
}

func (s service) Get(id string) (*entities.Bank, error) {
	var bank *entities.Bank
	tx := s.db.Find(&bank, "id = ?", id)
	if err := tx.Error; err != nil {
		return nil, err
	}

	return bank, nil
}

func (s service) Insert(bank *entities.Bank) error {
	tx := s.db.Create(bank)
	err := tx.Error
	return err
}

func (s service) Update(bank *entities.Bank) error {
	tx := s.db.Update(bank.ID, bank)
	err := tx.Error
	return err
}

func (s service) Delete(id string) error {
	tx := s.db.Delete(&entities.Bank{}, id)
	err := tx.Error
	return err
}
