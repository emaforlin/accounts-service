package main

import (
	"log"

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
		Name:  "accounts-service",
		Level: hclog.Info,
	}), conf, db).Start()
}
