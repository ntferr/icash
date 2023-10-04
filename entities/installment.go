package entities

type Installment struct {
	ID     string `json:"id"`
	CardID string `json:"card_id"`
	Number int    `json:"number"`
}
