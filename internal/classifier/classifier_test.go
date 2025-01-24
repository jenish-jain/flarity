package classifier_test

import (
	"testing"

	"github.com/jenish-jain/flarity/internal/classifier"
	"github.com/stretchr/testify/suite"
)

type ClassifierTestSuite struct {
	suite.Suite
	classifier classifier.Classifier
}

func (s *ClassifierTestSuite) SetupTest() {
	s.classifier = classifier.NewClassifier()
}

func (s *ClassifierTestSuite) TestClassify() {
	s.Equal("food", s.classifier.Classify("restaurant"))
	s.Equal("travel", s.classifier.Classify("flight"))
	s.Equal("unknown", s.classifier.Classify(""))
	s.Equal("clothes", s.classifier.Classify("zara"))
}

func (s *ClassifierTestSuite) TestClassifyCaseInsensitive() {
	s.Equal("food", s.classifier.Classify("RESTAURANT"))
	s.Equal("travel", s.classifier.Classify("FLIGHT"))
	s.Equal("unknown", s.classifier.Classify(""))
	s.Equal("clothes", s.classifier.Classify("ZarA"))
}

func (s *ClassifierTestSuite) TestClassifyWithMultipleKeywords() {
	s.Equal("food", s.classifier.Classify("Avighna Fast Foods"))
	s.Equal("travel", s.classifier.Classify("rapido"))
}

func TestClassierTestSuite(t *testing.T) {
	suite.Run(t, new(ClassifierTestSuite))
}
