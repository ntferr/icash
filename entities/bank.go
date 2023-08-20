package entities

import "errors"

type Bank struct {
	ID   string `gorm:"primaryKey"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func (b Bank) Validate() error {
	if b.Name == "" {
		return errors.New("bank name is required")
	}
	if b.Code == "" {
		return errors.New("bank code is required")
	}
	if len(b.Code) != 3 {
		return errors.New("insert a bank code valid")
	}

	return nil
}
