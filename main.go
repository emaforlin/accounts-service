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
		// Level: hclog.Debug,
	})

	logger.Info("Load configurations")
	config.InitViper("config")
	conf := config.LoadConfig()
	logger.Debug("Permissions loaded", conf.AccessControl)

	logger.Info("Connect to database")
	db := database.NewMySQLDatabase(conf, logger)

	if conf.Db.Migrate {
		logger.Info("Migrating database...")
		err := db.AutoMigrate()
		if err != nil {
			logger.Error(err.Error())
			logger.Info("Trying to run server anyway")
		} else {
			logger.Info("Migration done!")
		}
	}

	logger.Info("Setup server...")
	server.NewRPCServer(logger, *conf, db).Start()
}
