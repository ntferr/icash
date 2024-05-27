package entities

import "errors"

type Card struct {
	ID       string `json:"id" gorm:"primaryKey"`
	BankID   string `json:"bank_id"`
	Number   string `json:"number"`
	ExpireAt string `json:"expire_at"`
	Debts    []Debt `json:"debts"  gorm:"foreignKey:CardID"`
}

func (card Card) Validate() error {
	if err := numberValidate(card.Number); err != nil {
		return err
	}
	return nil
}

func numberValidate(number string) error {
	if number == "" {
		return errors.New("number must have a value")
	}
	return nil
}
