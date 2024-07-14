package main

import (
	"log"

	"github.com/emaforlin/accounts-service/config"
	"github.com/emaforlin/accounts-service/database"
	"github.com/emaforlin/accounts-service/server"
	"github.com/hashicorp/go-hclog"
)

func main() {
	logger := hclog.FromStandardLogger(log.Default(), &hclog.LoggerOptions{
		Name:  "accounts-service",
		Level: hclog.Info,
	})

	logger.Info("Load configurations")
	config.InitViper("config.yaml")
	conf := config.LoadConfig()

	logger.Info("Connect to database")
	db := database.NewMySQLDatabase(conf)

	logger.Info("Setup server...")
	server.NewRPCServer(logger, conf, db).Start()
}
