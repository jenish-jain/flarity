package config_test

import (
	"testing"

	"github.com/jenish-jain/flarity/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
	config *config.Config
}

func (s *ConfigTestSuite) SetupTest() {
	s.config = config.InitConfig("test")
}

func (s *ConfigTestSuite) TestGetLogLevel() {
	assert.Equal(s.T(), "debug", s.config.GetLogLevel())
}

func (s *ConfigTestSuite) TestGetServerPort() {
	assert.Equal(s.T(), "8080", s.config.GetServerPort())
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
