package transaction

import "time"

// TransactionType represents the type of transaction (credit/debit)
type TransactionType string

const (
	Credit TransactionType = "Credit"
	Debit  TransactionType = "Debit"
)

type Transaction struct {
	ID       string          `json:"_id" bson:"_id"`
	Date     time.Time       `json:"date" bson:"date"`
	Amount   float64         `json:"amount" bson:"amount"`
	Type     TransactionType `json:"type" bson:"type"`
	Currency string          `json:"currency" bson:"currency"`
	Meta     Meta            `json:"meta" bson:"meta"`
}

type Meta struct {
	ClientTxnID string `json:"client_txn_id" bson:"client_txn_id"`
	ClientName  string `json:"client_name" bson:"client_name"`
	Description string `json:"description" bson:"description"`
	Status      string `json:"status" bson:"status"`
	Category    string `json:"category" bson:"category"`
}

type Amount float64
