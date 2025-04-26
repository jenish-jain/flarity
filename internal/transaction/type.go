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

type TransactionSummary struct {
	Category    string          `json:"category" bson:"category"`
	TotalAmount float64         `json:"total_amount" bson:"total_amount"`
	Type        TransactionType `json:"type" bson:"type"`
}
type TransactionSummaryResponse struct {
	Year          int                  `json:"year" bson:"year"`
	Month         int                  `json:"month" bson:"month"`
	TotalDebit    float64              `json:"total_debit" bson:"total_debit"`
	TotalCredit   float64              `json:"total_credit" bson:"total_credit"`
	CategorySplit []TransactionSummary `json:"category_split" bson:"category_split"`
}

func (tt *TransactionType) IsCredit() bool {
	return *tt == Credit
}

func (tt *TransactionType) IsDebit() bool {
	return *tt == Debit
}
