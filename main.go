package main

import (
	"github.com/emaforlin/accounts-service/config"
	"github.com/emaforlin/accounts-service/database"
	"github.com/emaforlin/accounts-service/server"
	"github.com/hashicorp/go-hclog"
)

func main() {
	config.InitViper("config.yaml")
	conf := config.LoadConfig()
	db := database.NewMySQLDatabase(conf)

	go server.NewRPCServer(hclog.Default(), conf, db).Start()
	server.NewEchoServer(conf, db).Start()

	select {}
}
