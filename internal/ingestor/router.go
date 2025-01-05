package ingestor

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes(router *gin.Engine) {
	fileGroup := router.Group("ingest")
	fileGroup.POST("/takeout", h.IngestTakeoutRecords)
}
