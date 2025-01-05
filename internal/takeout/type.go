package takeout

type Takeout struct {
	Transactions []Record `json:"transactions"`
}

type Record struct {
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	Title         string  `json:"title"`
	Account       *string `json:"account"`
	Time          string  `json:"time"`
	Product       string  `json:"product"`
	TransactionID string  `json:"transactionId"`
	Status        string  `json:"status"`
}
