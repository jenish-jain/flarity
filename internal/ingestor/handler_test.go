package ingestor_test

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jenish-jain/flarity/internal/config"
	"github.com/jenish-jain/flarity/internal/ingestor"
	"github.com/jenish-jain/logger"
	"github.com/stretchr/testify/suite"
)

type IngestorTestSuite struct {
	suite.Suite
	config *config.Config
	router *gin.Engine
}

func (s *IngestorTestSuite) SetupTest() {
	s.config = config.InitConfig("test")
	logger.Init(s.config.LogLevel)
	handler := ingestor.NewHandler()
	s.router = gin.Default()
	s.router.POST("/ingest/takeout", handler.IngestTakeoutRecords)
}

func (s *IngestorTestSuite) TestIngestTakeoutRecordsForEmptyFile() {
	req, _ := http.NewRequest("POST", "/ingest/takeout", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *IngestorTestSuite) TestIngestTakeoutRecordsSuccess() {
	// Create a sample JSON file
	fileContent := `[{
        "currency": "₹",
        "amount": 10,
        "title": "instamojo",
        "account": "XXXX068726",
        "time": "04-11-2019",
        "product": "Google Pay",
        "transactionId": "ICIbb6d95a1e7f54f2a907f38bd82f15e2a",
        "status": "Completed"
    }]`
	file, err := os.CreateTemp("", "sample*.json")
	s.Require().NoError(err)
	defer os.Remove(file.Name())

	_, err = file.Write([]byte(fileContent))
	s.Require().NoError(err)
	file.Close()

	// Create a multipart form request with the JSON file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	s.Require().NoError(err)

	file, err = os.Open(file.Name())
	s.Require().NoError(err)
	defer file.Close()

	_, err = io.Copy(part, file)
	s.Require().NoError(err)
	writer.Close()

	req, err := http.NewRequest("POST", "/ingest/takeout", body)
	s.Require().NoError(err)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)

}

func (s *IngestorTestSuite) TearDownTest() {
	err := os.RemoveAll(filepath.Dir(s.config.AssetsPath))
	if err != nil {
		log.Fatal(err)
	}

}

func TestIngestorTestSuite(t *testing.T) {
	suite.Run(t, new(IngestorTestSuite))
}
