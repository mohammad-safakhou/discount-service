package transport

import (
	"database/sql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type ApplicationContext struct {
	VConfig *viper.Viper
	MysqlDb  *sql.DB
	Logger  *zap.SugaredLogger
}