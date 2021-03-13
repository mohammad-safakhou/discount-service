package database

import (
	"database/sql"
	"github.com/spf13/viper"
)

func mysqlConnectionString(config *viper.Viper) string {
	return config.GetString("database.mysql.uri")
}

func (ac *DBContext) MysqlConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", mysqlConnectionString(ac.VConfig))
	if err != nil {
		return nil, err
	}
	return db, nil
}