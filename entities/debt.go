package entities

type Debt struct {
	ID          int `json:"id" gorm:"primaryKey"`
	CardRefer   int
	TicketRefer int
	Value       float32 `json:"value"`
	Recurrent   bool    `json:"recurrent"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	DueDate     string  `json:"due_date"`
	Card        Card    `json:"card" gorm:"foreignKey:CardRefer"`
	Ticket      Ticket  `json:"ticket" gorm:"foreignKey:TicketRefer"`
}
