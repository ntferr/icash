package entities

type Ticket struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	CodeBars string `json:"code_bars"`
}
