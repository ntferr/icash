package bank

import (
	"errors"

	"github.com/ntferr/icash/entities"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

type Contract interface {
	FindAll() (*[]entities.Bank, error)
	FindByID(bankID *string) (*entities.Bank, error)
	Insert(bank *entities.Bank) error
	Update(bank *entities.Bank) error
	Delete(bankID *string) error
}

func NewBankCRUD(db *gorm.DB) service {
	return service{
		db: db,
	}
}

func (s service) FindAll() (*[]entities.Bank, error) {
	var banks []entities.Bank
	err := s.db.
		Preload("Cards").
		Find(&banks).Error
	return &banks, err
}

func (s service) FindByID(bankID *string) (*entities.Bank, error) {
	var bank entities.Bank
	err := s.db.First(bank).Where("id = ?", bankID).Error
	return &bank, err
}

func (s service) Insert(bank *entities.Bank) error { return s.db.Create(bank).Error }

func (s service) Update(bank *entities.Bank) error {
	tx := s.db.Updates(bank)
	if tx.RowsAffected == 0 {
		return errors.New("id doesn't exist")
	}
	return tx.Error
}

func (s service) Delete(bankID *string) error {
	tx := s.db.Delete(&entities.Bank{}).Where("id = ?", bankID)
	if tx.RowsAffected == 0 {
		return errors.New("id doesn't exist")
	}
	return tx.Error
}
