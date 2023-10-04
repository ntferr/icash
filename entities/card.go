package entities

type Card struct {
	ID       string `json:"id" gorm:"primaryKey"`
	BankID   string `json:"bank_id"`
	Number   string `json:"number"`
	ExpireAt string `json:"expire_at"`
	Debts    []Debt `json:"debts"  gorm:"foreignKey:CardID"`
}

func (card Card) Validate() {

}
