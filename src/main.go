package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/valverde.thiago/go-agenda-api/api"
	"github.com/valverde.thiago/go-agenda-api/contact"
	"github.com/valverde.thiago/go-agenda-api/metrics"
	"github.com/valverde.thiago/go-agenda-api/util"
	mgo "gopkg.in/mgo.v2"
)

var database *mgo.Database

func main() {
	router := gin.Default()
	metricService := metrics.NewMetricService("/metrics", router)
	metricService.Configure()
	config := util.LoadEnvConfig("./.", "app")
	database = util.ConnectToDatabase(config)
	store := contact.NewMongoDbStore(database)
	server := api.NewServer(router, &config)
	contactController := contact.NewController(store)
	server.ConfigureController(contactController)

	err := server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Failed to start HTTP server")
	}

}
