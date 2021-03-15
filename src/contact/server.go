package contact

import (
	"github.com/gin-gonic/gin"
	"github.com/valverde.thiago/go-agenda-api/metrics"
	"github.com/valverde.thiago/go-agenda-api/util"
)

// Server server
type Server struct {
	store  Store
	router *gin.Engine
}

// NewServer creates a new server instance
func NewServer(store Store, router *gin.Engine, config util.Config) *Server {
	server := &Server{
		store:  store,
		router: router,
	}
	contactController := NewController(store)
	contactController.SetupRoutes(server.router)
	metricsController := metrics.NewController()
	metricsController.SetupRoutes(server.router)
	return server
}

// Start runs the HTTP Server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
