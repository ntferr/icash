package entities

import "time"

type Debt struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	CardID      string    `json:"card_id"`
	TicketID    string    `json:"ticket_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Value       float32   `json:"value"`
	Recurrent   bool      `json:"recurrent"`
	DueDate     time.Time `json:"due_date"`
	Card        Card      `json:"card" gorm:"foreignKey:CardID"`
	Ticket      Ticket    `json:"ticket" gorm:"foreignKey:TicketID"`
}
