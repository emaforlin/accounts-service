package main

import (
	"github.com/emaforlin/accounts-service/config"
	"github.com/emaforlin/accounts-service/database"
	"github.com/emaforlin/accounts-service/x/entities"
	"github.com/hashicorp/go-hclog"
)

func main() {
	config.InitViper("config.yaml")
	cfg := config.LoadConfig()
	db := database.NewMySQLDatabase(cfg, hclog.Default())
	db.GetDb().AutoMigrate(&entities.InsertFoodPlaceDto{})
	db.GetDb().AutoMigrate(&entities.InsertPersonDto{})
}
