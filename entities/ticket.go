package entities

type Ticket struct {
	Id       int    `json:"id"`
	CodeBars string `json:"code_bars"`
	Debt     Debt   `json:"debt"`
}
