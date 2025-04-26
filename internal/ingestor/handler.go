package ingestor

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jenish-jain/flarity/internal/config"
	"github.com/jenish-jain/flarity/internal/takeout"
	"github.com/jenish-jain/flarity/internal/transaction"
	"github.com/jenish-jain/flarity/pkg/files"
	"github.com/jenish-jain/logger"
)

type Handler struct {
	config          *config.Config
	takeoutService  takeout.Service
	transactionRepo transaction.Repository
}

func (h *Handler) IngestTakeoutRecords(ctx *gin.Context) {
	logger.Debug("Ingesting takeout records")

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		logger.ErrorWithCtx(ctx, "no file found | bad request", "err", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileBytes, err := getFileBytesFromMultipartHeaders(fileHeader)

	if err != nil {
		logger.ErrorWithCtx(ctx, "error reading file", "err", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	takeout := h.takeoutService.
		Get(fileBytes)
	transactions := takeout.ToTransactions()

	files.Write(h.config.AssetsPath, transactions)
	//  for each transaction insert a record into tractions collection via transaction repo
	for _, transaction := range transactions {
		if transaction.Amount != 0 && transaction.Meta.Status == "Completed" {
			if err := h.transactionRepo.Add(&transaction); err != nil {
				logger.ErrorWithCtx(ctx, "error adding transaction", "err", err)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
	}

	logger.InfoWithCtx(ctx, "All transactions processed successfully")
	ctx.JSON(http.StatusOK, gin.H{"message": "All transactions processed successfully"})
}

func getFileBytesFromMultipartHeaders(fileHeader *multipart.FileHeader) ([]byte, error) {
	takeoutFile, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer takeoutFile.Close()

	//Create temp file
	tempFile, err := os.CreateTemp("", fileHeader.Filename)
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()

	//Delete temp file after importing
	defer os.Remove(tempFile.Name())

	fileBytes, err := io.ReadAll(takeoutFile)
	if err != nil {
		return nil, err
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

func NewHandler(transactionRepo transaction.Repository) *Handler {
	return &Handler{
		config:          config.GetConfig(),
		takeoutService:  takeout.NewService(),
		transactionRepo: transactionRepo,
	}
}
