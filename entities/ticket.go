package entities

type Ticket struct {
	Id       string `json:"id" gorm:"primaryKey"`
	CodeBars string `json:"code_bars"`
}
