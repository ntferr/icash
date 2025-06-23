package debt

import (
	"errors"

	"github.com/icash/internal/controller/ticket"
)

type Debt struct {
	ID           string        `json:"id" gorm:"primaryKey"`
	CardID       string        `json:"card_id"`
	TicketID     string        `json:"ticket_id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Ticket       ticket.Ticket `json:"ticket" gorm:"foreignKey:DebtID"`
	Installments []Installment `json:"installments" gorm:"foreignKey:DebtID"`
}

type Installment struct {
	ID      string `json:"id" gorm:"primaryKey"`
	DebtID  string `json:"debt_id"`
	DueDate string `json:"due_date" gorm:"dueDate"`
	Paid    bool   `json:"paid" gorm:"paid"`
	Number  int    `json:"number" gorm:"number"`
}

func (debt Debt) Validate() error {
	if debt.Name == "" {
		return errors.New("debt must have a name")
	}
	return nil
}
