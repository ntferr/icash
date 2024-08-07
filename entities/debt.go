package entities

import "errors"

type Debt struct {
	ID           string        `json:"id" gorm:"primaryKey"`
	CardID       string        `json:"card_id"`
	TicketID     string        `json:"ticket_id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Ticket       Ticket        `json:"ticket" gorm:"foreignKey:DebtID"`
	Installments []Installment `json:"installments" gorm:"foreignKey:DebtID"`
}

func (debt Debt) Validate() error {
	if debt.Name == "" {
		return errors.New("debt must have a name")
	}
	return nil
}
