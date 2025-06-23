package ticket

type Ticket struct {
	ID       string `json:"id" gorm:"primaryKey"`
	DebtID   string `json:"debt_id"`
	CodeBars string `json:"code_bars"`
}
