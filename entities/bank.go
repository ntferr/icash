package entities

import "errors"

type Bank struct {
	ID     string `gorm:"primaryKey"`
	Name   string `json:"name"`
	Number string `json:"number"`
}

func (b Bank) Validate() error {
	if b.Name == "" {
		err := errors.New("bank name is required")
		return err
	}
	if b.Number == "" {
		err := errors.New("bank number is required")
		return err
	}

	return nil
}
