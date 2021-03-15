package contact

import (
	"github.com/gin-gonic/gin"
)

// Server server
type Server struct {
	store  Store
	router *gin.Engine
}

// NewServer creates a new server instance
func NewServer(store Store) *Server {
	server := &Server{
		store:  store,
		router: gin.Default(),
	}
	NewController(store, server.router)
	return server
}

// Start runs the HTTP Server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
