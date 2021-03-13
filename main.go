package main

import (
	"log"

	"github.com/valverde.thiago/go-agenda-api/api"
	"github.com/valverde.thiago/go-agenda-api/util"
	mgo "gopkg.in/mgo.v2"
)

var database *mgo.Database

func main() {
	config := util.LoadEnvConfig("./.", "app")
	database = util.ConnectToDatabase(config)
	server := api.NewServer(database)
	err := server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Failed to start HTTP server")
	}

}
