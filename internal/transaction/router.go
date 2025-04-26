package transaction

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes(router *gin.Engine) {
	fileGroup := router.Group("transactions")
	fileGroup.GET("", h.GetTransactions)
	fileGroup.GET("/summary", h.GetMonthlyTransactionSummary)
}
