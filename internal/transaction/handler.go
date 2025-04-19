package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/jenish-jain/flarity/internal/config"
	"go.mongodb.org/mongo-driver/bson"
)

type Handler struct {
	config     *config.Config
	repository Repository
}

func (h *Handler) GetTransactions(ctx *gin.Context) {
	page := 1
	limit := 10
	transactions, err := h.repository.GetByFilter(bson.D{}, page, limit)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, transactions)
}

func NewHandler(config *config.Config, repository Repository) *Handler {
	return &Handler{
		config:     config,
		repository: repository,
	}
}
