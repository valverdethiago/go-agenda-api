package api

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
	Router *gin.Engine
	config *config.Config
}

// Controller controller
type Controller interface {
	SetupRoutes(router *gin.Engine)
}

// NewServer creates a new server instance
func NewServer(router *gin.Engine, config *config.Config) *Server {
	server := &Server{
		Router: router,
		config: config,
	}
	server.ConfigureLogging()
	return server
}

func (server *Server) ConfigureController(controller Controller) {
	controller.SetupRoutes(server.Router)
}

// ConfigureLogging configure gin logs
func (server *Server) ConfigureLogging() {
	flag.Parse()
	server.Router.Use(ginglog.Logger(3 * time.Second))
	server.Router.Use(gin.Recovery())
}

// Start runs the HTTP Server on a specific address
func (server *Server) Start(address string) error {
	var readTimeout time.Duration
	readTimeout = time.Duration(server.config.ReadTimeout) * time.Second
	var writeTimeout time.Duration
	writeTimeout = time.Duration(server.config.WriteTimeout) * time.Second

	s := &http.Server{
		Addr:         server.config.ServerAddress,
		Handler:      server.Router,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
	return s.ListenAndServe()
}
