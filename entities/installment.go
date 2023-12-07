package entities

type Installment struct {
	ID      string `json:"id" gorm:"primaryKey"`
	DueDate string `json:"due_date" gorm:"dueDate"`
	Paid    bool   `json:"paid" gorm:"paid"`
	Number  int    `json:"number" gorm:"number"`
}
