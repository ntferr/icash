package entities

import (
	"errors"
)

type Bank struct {
	ID    string `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Code  string `json:"code"`
	Cards []Card `json:"cards" gorm:"foreignKey:BankID"`
}

func (bank Bank) Validate() error {
	if bank.Name == "" {
		return errors.New("bank name is required")
	}
	if err := codeValidate(bank.Code); err != nil {
		return err
	}
	return nil
}

func codeValidate(code string) error {
	if code == "" {
		return errors.New("bank code is required")
	}
	if len(code) != 3 {
		return errors.New("insert a valid code bank")
	}
	return nil
}
