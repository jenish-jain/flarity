package takeout_test

import (
	"testing"

	"github.com/jenish-jain/flarity/internal/takeout"
	"github.com/jenish-jain/flarity/internal/transaction"
	"github.com/stretchr/testify/suite"
)

type TakeoutTestSuite struct {
	suite.Suite
	service takeout.Service
}

func (s *TakeoutTestSuite) SetupTest() {
	s.service = takeout.NewService()
}

func (s *TakeoutTestSuite) TestToTransactions() {
	takeout := takeout.Takeout{
		Transactions: []takeout.Record{
			{
				Currency:      "INR",
				Amount:        100,
				Title:         "Title",
				Account:       nil,
				Time:          "06-10-2024",
				Product:       "Google Pay",
				TransactionID: "TransactionID",
				Status:        "Status",
			},
		},
	}
	transactions := takeout.ToTransactions()
	s.Equal(1, len(transactions))
	s.Equal("INR", transactions[0].Currency)
	s.Equal(100.0, transactions[0].Amount)
	s.Equal("Google Pay", transactions[0].Meta.ClientName)
	s.Equal("Status", transactions[0].Meta.Status)
	s.Equal("TransactionID", transactions[0].Meta.ClientTxnID)
	s.Equal("Title", transactions[0].Meta.Description)
	s.Equal("06-10-2024", transactions[0].Date.Format("02-01-2006"))
	s.Equal(transaction.Credit, transactions[0].Type)

}

func TestTakeoutTestSuite(t *testing.T) {
	suite.Run(t, new(TakeoutTestSuite))
}
