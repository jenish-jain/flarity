package server

import (
	"github.com/gin-contrib/cors"
	"github.com/jenish-jain/flarity/internal/ingestor"
	"github.com/jenish-jain/flarity/internal/login"
	"github.com/jenish-jain/logger"
)

type Handlers struct {
	ingestorHandler *ingestor.Handler
	loginHandler    *login.Handler
}

func (s *Server) InitRoutes(h Handlers) {
	router := s.routerGroups.rootRouter

	// Enable CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // Replace with your frontend's origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.Use(logger.AttachRequestIdToRequests)

	h.ingestorHandler.InitRoutes(router)
	h.loginHandler.InitRoutes(router.Group(""))
}

func InitHandlers(ingestorHandler *ingestor.Handler, loginHandler *login.Handler) Handlers {
	return Handlers{
		ingestorHandler: ingestorHandler,
		loginHandler:    loginHandler,
	}
}
