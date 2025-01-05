package takeout_test

func (s *TakeoutTestSuite) TestGetSuccess() {
	fileBytes := `[{"currency":"₹","amount":10,"title":"instamojo","account":"XXXX068726","time":"04-11-2019","product":"Google Pay","transactionId":"ICIbb6d95a1e7f54f2a907f38bd82f15e2a","status":"Completed"}]`
	takeout := s.service.Get([]byte(fileBytes))
	s.Equal(1, len(takeout.Transactions))
	s.Equal("₹", takeout.Transactions[0].Currency)
	s.Equal(10.0, takeout.Transactions[0].Amount)
	s.Equal("instamojo", takeout.Transactions[0].Title)
	s.Equal("XXXX068726", *takeout.Transactions[0].Account)
	s.Equal("04-11-2019", takeout.Transactions[0].Time)
	s.Equal("Google Pay", takeout.Transactions[0].Product)
	s.Equal("ICIbb6d95a1e7f54f2a907f38bd82f15e2a", takeout.Transactions[0].TransactionID)
	s.Equal("Completed", takeout.Transactions[0].Status)
}
