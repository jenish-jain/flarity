package takeout

import (
	"github.com/google/uuid"
	"github.com/jenish-jain/flarity/internal/classifier"
	"github.com/jenish-jain/flarity/internal/transaction"
	"github.com/jenish-jain/flarity/pkg/datetime"
)

func (t *Takeout) ToTransactions() []transaction.Transaction {
	var transactions []transaction.Transaction
	classifier := classifier.NewClassifier()

	for _, takeoutTxn := range t.Transactions {
		date, _ := datetime.StringToDate(takeoutTxn.Time, datetime.DDMMYYYY)
		transactions = append(transactions, transaction.Transaction{
			ID:       uuid.NewString(),
			Currency: takeoutTxn.Currency,
			Amount:   takeoutTxn.Amount,
			Date:     date,
			Type:     transaction.GetTransactionType(takeoutTxn.Amount),
			Meta: transaction.Meta{
				ClientTxnID: takeoutTxn.TransactionID,
				ClientName:  takeoutTxn.Product,
				Description: takeoutTxn.Title,
				Status:      takeoutTxn.Status,
				Category:    classifier.Classify(takeoutTxn.Title),
			},
		})
	}
	return transactions
}
