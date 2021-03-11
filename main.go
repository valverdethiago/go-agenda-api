package main

import (
	"log"

	"github.com/valverde.thiago/go-agenda-api/api"
	"github.com/valverde.thiago/go-agenda-api/util"
	mgo "gopkg.in/mgo.v2"
)

var database *mgo.Database

func main() {
	config := loadConfig()
	database = connectToDatabase(config)
	server := api.NewServer(database)
	err := server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Failed to start HTTP server")
	}

}

func loadConfig() util.Config {
	config, err := util.LoadConfig(".", "app")
	if err != nil {
		log.Fatal("Error loading application config: ", err)
	}
	return config
}

func connectToDatabase(config util.Config) *mgo.Database {
	session, err := mgo.Dial(config.DBServer)
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}
	return session.DB(config.DBName)
}
