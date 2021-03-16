package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/valverde.thiago/go-agenda-api/contact"
	"github.com/valverde.thiago/go-agenda-api/metrics"
	"github.com/valverde.thiago/go-agenda-api/util"
	mgo "gopkg.in/mgo.v2"
)

var database *mgo.Database

func main() {
	router := gin.Default()
	config := util.LoadEnvConfig("./.", "app")
	database = util.ConnectToDatabase(config)
	store := contact.NewMongoDbStore(database)
	server := contact.NewServer(store, router, &config)
	err := server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Failed to start HTTP server")
	}
	configureMiddleware(config, router)

}

func configureMiddleware(config util.Config, router *gin.Engine) {
	metricService, err := metrics.NewPrometheusService(config.PrometheusAddress)
	if err != nil {
		log.Fatal(err.Error())
	}
	router.Use(metrics.Metrics(metricService))
}
