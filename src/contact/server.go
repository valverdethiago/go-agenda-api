package contact

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/valverde.thiago/go-agenda-api/util"
)

// Server server
type Server struct {
	store  Store
	router *gin.Engine
	config *util.Config
}

// NewServer creates a new server instance
func NewServer(store Store, router *gin.Engine, config *util.Config) *Server {
	server := &Server{
		store:  store,
		router: router,
		config: config,
	}
	contactController := NewController(store)
	contactController.SetupRoutes(server.router)
	return server
}

// Start runs the HTTP Server on a specific address
func (server *Server) Start(address string) error {
	var readTimeout time.Duration
	readTimeout = time.Duration(server.config.ReadTimeout) * time.Second
	var writeTimeout time.Duration
	writeTimeout = time.Duration(server.config.WriteTimeout) * time.Second

	s := &http.Server{
		Addr:         server.config.ServerAddress,
		Handler:      server.router,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
	return s.ListenAndServe()
}
