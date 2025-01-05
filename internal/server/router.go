package server

import (
	"github.com/jenish-jain/flarity/internal/ingestor"
	"github.com/jenish-jain/logger"
)

func (s *Server) InitRoutes() {
	router := s.routerGroups.rootRouter
	router.Use(logger.AttachRequestIdToRequests)

	ingestorHandler := ingestor.NewHandler()
	ingestorHandler.InitRoutes(router)
}
