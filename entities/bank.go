package entities

import "errors"

type Bank struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func (bank Bank) Validate() error {
	if bank.Name == "" {
		return errors.New("bank name is required")
	}
	if bank.Code == "" {
		return errors.New("bank code is required")
	}
	if len(bank.Code) != 3 {
		return errors.New("insert a bank code valid")
	}

	return nil
}
