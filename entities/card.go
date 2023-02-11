package entities

type Card struct {
	ID        string `json:"id" gorm:"primaryKey"`
	BankRefer int
	Number    string `json:"number"`
	ExpireAt  string `json:"expire_at"`
	Bank      Bank   `json:"bank" gorm:"foreignKey:BankRefer"`
}
