package util

import (
	"log"

	"github.com/valverde.thiago/go-agenda-api/config"
	"gopkg.in/mgo.v2"
)

// LoadEnvConfig load app configuration based on file
func LoadEnvConfig(path string, file string) config.Config {
	config, err := config.LoadConfig(path, file)
	if err != nil {
		log.Fatal("Error loading application config: ", err)
	}
	return config
}

// ConnectToDatabase connect to mongo database
func ConnectToDatabase(config config.Config) *mgo.Database {
	session, err := mgo.Dial(config.DBServer)
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}
	return session.DB(config.DBName)
}
