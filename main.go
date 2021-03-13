package main

//go:generate sqlboiler --wipe mysql
import (
	"discount-service/logger"
	"discount-service/transport/rest"
)
func main() {
	l := logger.ZapLogger()
	s := l.Sugar()
	rest.StartRestServer(s)
}
