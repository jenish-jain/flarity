package transaction

import "time"

// TransactionType represents the type of transaction (credit/debit)
type TransactionType string

const (
	Credit TransactionType = "Credit"
	Debit  TransactionType = "Debit"
)

type Transaction struct {
	ID       string          `json:"_id"`
	Date     time.Time       `json:"date"`
	Amount   float64         `json:"amount"`
	Type     TransactionType `json:"type"`
	Currency string          `json:"currency"`
	Meta     Meta            `json:"meta"`
}

type Meta struct {
	ClientTxnID string `json:"client_txn_id"`
	ClientName  string `json:"client_name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Category    string `json:"category"`
}

type Amount float64
