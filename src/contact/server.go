package contact

import (
	"flag"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ginglog "github.com/szuecs/gin-glog"
	"github.com/valverde.thiago/go-agenda-api/config"
)

// Server server
type Server struct {
	store  Store
	router *gin.Engine
	config *config.Config
}

// NewServer creates a new server instance
func NewServer(store Store, router *gin.Engine, config *config.Config) *Server {
	server := &Server{
		store:  store,
		router: router,
		config: config,
	}
	contactController := NewController(store)
	contactController.SetupRoutes(server.router)
	server.ConfigureLogging()
	return server
}

// ConfigureLogging configure gin logs
func (server *Server) ConfigureLogging() {
	flag.Parse()
	server.router.Use(ginglog.Logger(3 * time.Second))
	server.router.Use(gin.Recovery())
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
