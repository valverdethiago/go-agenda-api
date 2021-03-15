package util

import (
	"log"

	"gopkg.in/mgo.v2"
)

// LoadEnvConfig load app configuration based on file
func LoadEnvConfig(path string, file string) Config {
	config, err := LoadConfig(path, file)
	if err != nil {
		log.Fatal("Error loading application config: ", err)
	}
	return config
}

// ConnectToDatabase connect to mongo database
func ConnectToDatabase(config Config) *mgo.Database {
	session, err := mgo.Dial(config.DBServer)
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}
	return session.DB(config.DBName)
}
