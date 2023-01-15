package entities

type Bank struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Number string `json:"number"`
}
