package ingestor

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jenish-jain/flarity/internal/config"
	"github.com/jenish-jain/flarity/internal/takeout"
	"github.com/jenish-jain/flarity/pkg/files"
	"github.com/jenish-jain/logger"
)

type Handler struct {
	config         *config.Config
	takeoutService takeout.Service
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

func NewHandler() *Handler {
	return &Handler{
		config:         config.GetConfig(),
		takeoutService: takeout.NewService(),
	}
}
