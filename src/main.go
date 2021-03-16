package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/valverde.thiago/go-agenda-api/contact"
	"github.com/valverde.thiago/go-agenda-api/util"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	mgo "gopkg.in/mgo.v2"
)

var database *mgo.Database

func main() {
	router := gin.New()

	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	config := util.LoadEnvConfig("./.", "app")
	database = util.ConnectToDatabase(config)
	store := contact.NewMongoDbStore(database)
	server := contact.NewServer(store, router, &config)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello world!")
	})

	err := server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Failed to start HTTP server")
	}

}
