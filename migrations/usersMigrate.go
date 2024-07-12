package main

import (
	"github.com/emaforlin/accounts-service/config"
	"github.com/emaforlin/accounts-service/database"
	"github.com/emaforlin/accounts-service/x/entities"
)

func main() {
	config.InitViper("config.yaml")
	cfg := config.LoadConfig()
	db := database.NewMySQLDatabase(cfg)
	db.GetDb().AutoMigrate(&entities.InsertFoodPlaceDto{})
}
