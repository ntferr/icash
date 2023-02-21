package entities

type Bank struct {
	ID     string `gorm:"primaryKey"`
	Name   string `json:"name"`
	Number string `json:"number"`
}
