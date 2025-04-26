package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jenish-jain/flarity/internal/config"
	"go.mongodb.org/mongo-driver/bson"
)

type Handler struct {
	config     *config.Config
	repository Repository
}

func (h *Handler) GetTransactions(ctx *gin.Context) {
	// Extract and validate year and month using the common validator
	year, month, err := ValidateYearAndMonth(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Build the filter for MongoDB query
	filter := bson.M{}
	if year > 0 && month > 0 {
		filter = bson.M{
			"$expr": bson.M{
				"$and": []bson.M{
					{"$eq": []interface{}{bson.M{"$year": "$date"}, year}},
					{"$eq": []interface{}{bson.M{"$month": "$date"}, month}},
				},
			},
		}
	}

	// Pagination parameters
	page := 1
	limit := 10

	// Fetch transactions from the repository
	transactions, err := h.repository.GetByFilter(filter, page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the transactions
	ctx.JSON(http.StatusOK, transactions)
}

func (h *Handler) GetMonthlyTransactionSummary(ctx *gin.Context) {
	// Validate year and month using the common validator
	year, month, err := ValidateYearAndMonth(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Build the aggregation pipeline
	pipeline := bson.A{
		bson.M{
			"$match": bson.M{
				"meta.status": "Completed",
				"amount":      bson.M{"$gt": 0},
			},
		},
		bson.M{
			"$addFields": bson.M{
				"month": bson.M{"$month": "$date"},
				"year":  bson.M{"$year": "$date"},
			},
		},
		bson.M{
			"$match": bson.M{
				"month": month,
				"year":  year,
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"type":     "$type",
					"category": "$meta.category",
				},
				"totalAmount": bson.M{"$sum": "$amount"},
			},
		},
		bson.M{
			"$sort": bson.M{
				"_id.type":     1,
				"_id.category": 1,
			},
		},
	}

	// Execute the aggregation query
	rawResults, err := h.repository.GetAggregatedByFilter(pipeline)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map raw results to TransactionSummary and calculate totals
	var summaries []TransactionSummary
	var totalDebit, totalCredit float64

	for _, raw := range rawResults {
		group := raw["_id"].(bson.M)
		transactionType := TransactionType(group["type"].(string))
		totalAmount := raw["totalAmount"].(float64)

		// Add to the appropriate total
		if transactionType.IsDebit() {
			totalDebit += totalAmount
		} else if transactionType.IsCredit() {
			totalCredit += totalAmount
		}

		// Append to summaries
		summaries = append(summaries, TransactionSummary{
			Type:        transactionType,
			Category:    group["category"].(string),
			TotalAmount: totalAmount,
		})
	}

	// Create the response
	response := TransactionSummaryResponse{
		Year:          year,
		Month:         month,
		TotalDebit:    totalDebit,
		TotalCredit:   totalCredit,
		CategorySplit: summaries,
	}

	// Return the result
	ctx.JSON(http.StatusOK, response)
}

func NewHandler(config *config.Config, repository Repository) *Handler {
	return &Handler{
		config:     config,
		repository: repository,
	}
}
