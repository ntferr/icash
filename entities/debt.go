package entities

type Debt struct {
	ID           string        `json:"id" gorm:"primaryKey"`
	CardID       string        `json:"card_id"`
	TicketID     string        `json:"ticket_id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Recurrent    bool          `json:"recurrent"`
	Ticket       Ticket        `json:"ticket" gorm:"foreignKey:TicketID"`
	Installments []Installment `json:"installments" gorm:"foreignKey:CardID"`
}
