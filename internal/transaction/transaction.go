package transaction

func (t *Transaction) IsCredit() bool {
	return t.Type == Credit
}

func (t *Transaction) IsDebit() bool {
	return t.Type == Debit
}

func GetTransactionType(a float64) TransactionType {
	if a > 0 {
		return Credit
	}
	return Debit
}
