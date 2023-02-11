package entities

type Bank struct {
	ID     string `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Number string `json:"number"`
}
