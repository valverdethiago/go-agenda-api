package api

import (
	"flag"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	ginglog "github.com/szuecs/gin-glog"
	"github.com/valverde.thiago/go-agenda-api/contact"
	"github.com/valverde.thiago/go-agenda-api/util"
)

// Server server
type Server struct {
	store  contact.Store
	router *gin.Engine
	config *util.Config
}

// NewServer creates a new server instance
func NewServer(store contact.Store, router *gin.Engine, config *util.Config) *Server {
	server := &Server{
		store:  store,
		router: router,
		config: config,
	}
	contactController := contact.NewController(store)
	contactController.SetupRoutes(server.router)
	server.ConfigureLogging()
	return server
}

// ConfigureLogging configure gin logs
func (server *Server) ConfigureLogging() {
	flag.Parse()
	server.router.Use(ginglog.Logger(3 * time.Second))
	server.router.Use(gin.Recovery())
	glog.Warning("warning")
	glog.Error("err")
	glog.Info("info")
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
