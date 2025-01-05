package datetime_test

import (
	"testing"
	"time"

	"github.com/jenish-jain/flarity/pkg/datetime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DatetimeTestSuite struct {
	suite.Suite
}

func (s *DatetimeTestSuite) SetupTest() {
}

func (s *DatetimeTestSuite) TestStringToDateWithDDMMYYYYFormat() {
	date, err := datetime.StringToDate("06-03-1997", datetime.DDMMYYYY)
	assert.Equal(s.T(),
		time.Date(1997, time.March, 6, 0, 0, 0, 0, time.UTC),
		date)
	assert.Nil(s.T(), err)
}

func (s *DatetimeTestSuite) TestStringToDateWithInvalidFormat() {
	_, err := datetime.StringToDate("1997-03-06", datetime.DDMMYYYY)
	assert.NotNil(s.T(), err)
}

func TestDateTimeTestSuite(t *testing.T) {
	suite.Run(t, new(DatetimeTestSuite))
}
