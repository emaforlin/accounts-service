package main

import (
	"log"
	"time"

	"github.com/emaforlin/accounts-service/config"
	"github.com/emaforlin/accounts-service/database"
	"github.com/emaforlin/accounts-service/server"
	"github.com/hashicorp/go-hclog"
)

func main() {
	config.InitViper("config.yaml")
	conf := config.LoadConfig()
	db := database.NewMySQLDatabase(conf)
	server.NewRPCServer(hclog.FromStandardLogger(log.Default(), &hclog.LoggerOptions{
		Name:       "GRPC",
		Level:      hclog.Info,
		JSONFormat: true,
		TimeFormat: time.RFC3339,
	}), conf, db).Start()
}
