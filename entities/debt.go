package entities

type Debt struct {
	ID          int     `json:"id"`
	Type        int     `json:"type"`
	Value       float32 `json:"value"`
	Recurrent   bool    `json:"recurrent"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	DueDate     string  `json:"due_date"`
}
