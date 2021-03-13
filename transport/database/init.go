package database

import (
	"discount-service/transport"
	_ "github.com/go-sql-driver/mysql"
)

type DBContext struct {
	*transport.ApplicationContext
}

func (ac *DBContext) RegisterDatabases() {
	mysqlConnection, err := ac.MysqlConnect()
	if err != nil {
		ac.Logger.Fatal("error on connecting to mongo %v", err)
	}
	ac.MysqlDb = mysqlConnection
}
