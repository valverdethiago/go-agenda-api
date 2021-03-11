package api

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"

	contact "github.com/valverde.thiago/go-agenda-api/contact"
)

// Server server
type Server struct {
	database *mgo.Database
	router   *gin.Engine
}

// NewServer creates a new server instance
func NewServer(database *mgo.Database) Server {
	server := Server{
		database: database,
		router:   gin.Default(),
	}
	contactStore := contact.NewMongoDbStore(database)
	contact.NewController(contactStore, server.router)
	return server
}

// Start runs the HTTP Server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
