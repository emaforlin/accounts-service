package main

import (
	"github.com/emaforlin/accounts-service/config"
	"github.com/emaforlin/accounts-service/database"
	"github.com/emaforlin/accounts-service/x/entities"
	"github.com/hashicorp/go-hclog"
)

func main() {
	logger := hclog.New(hclog.DefaultOptions)
	config.InitViper("config.yaml")
	cfg := config.LoadConfig()
	db := database.NewMySQLDatabase(cfg, logger)

	db.GetDb().AutoMigrate(&entities.InsertFoodPlaceDto{})
	db.GetDb().AutoMigrate(&entities.InsertPersonDto{})
	logger.Info("Done!")
}
