package database

import (
	"discount-service/transport"
)

type DBContext struct {
	*transport.ApplicationContext
}

func (ac *DBContext) RegisterDatabases() {
	mongoClient, err := ac.MongoConnect()
	if err != nil {
		ac.Logger.Fatal("error on connecting to mongo %v", err)
	}
	ac.MongoDb = mongoClient
}
