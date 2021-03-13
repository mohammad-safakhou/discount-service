package transport

import (
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type ApplicationContext struct {
	VConfig *viper.Viper
	MongoDb  *mongo.Client
	Logger  *zap.SugaredLogger
}