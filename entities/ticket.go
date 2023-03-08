package entities

type Ticket struct {
	ID       string `json:"id" gorm:"primaryKey"`
	CodeBars string `json:"code_bars"`
}
