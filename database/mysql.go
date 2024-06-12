package database

import (
	"fmt"

	"github.com/emaforlin/accounts-service/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlDatabase struct {
	db *gorm.DB
}

func NewMySQLDatabase(cfg *config.Config) Database {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Db.User, cfg.Db.Passwd, cfg.Db.Host, cfg.Db.Name)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &mysqlDatabase{db: db}
}

func (p *mysqlDatabase) GetDb() *gorm.DB {
	return p.db
}
