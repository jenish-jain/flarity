package transaction_test

import (
	"testing"

	"github.com/jenish-jain/flarity/internal/transaction"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TransactionTestSuite struct {
	suite.Suite
}

func (s *TransactionTestSuite) SetupTest() {
}

func (s *TransactionTestSuite) TestIsCredit() {
	t := transaction.Transaction{
		Type: transaction.Credit,
	}
	assert.Equal(s.T(), true, t.IsCredit())
}

func (s *TransactionTestSuite) TestIsDebit() {
	t := transaction.Transaction{
		Type: transaction.Credit,
	}
	assert.Equal(s.T(), false, t.IsDebit())
}

func (s *TransactionTestSuite) TestGetTransactionType() {
	assert.Equal(s.T(), transaction.Credit, transaction.GetTransactionType(10))
	assert.Equal(s.T(), transaction.Debit, transaction.GetTransactionType(-10))
}

func TestTransactionTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionTestSuite))
}
