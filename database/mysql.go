package database

import (
	"errors"
	"fmt"

	"github.com/emaforlin/accounts-service/config"
	"github.com/emaforlin/accounts-service/x/entities"
	"github.com/hashicorp/go-hclog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type mysqlDatabase struct {
	db *gorm.DB
}

// AutoMigrate implements Database.
func (p *mysqlDatabase) AutoMigrate() error {
	var errorStr string
	err1 := p.GetDb().AutoMigrate(&entities.InsertPersonDto{})
	err2 := p.GetDb().AutoMigrate(&entities.InsertFoodPlaceDto{})

	errorStr = fmt.Sprintf("%v%v", err1, err2)
	if len(errorStr) != 0 {
		return errors.New("error migrating database: " + errorStr)
	}
	return nil
}

func NewMySQLDatabase(cfg *config.Config, l hclog.Logger) Database {
	dblogger := logger.New(l.Named("database").StandardLogger(&hclog.StandardLoggerOptions{}), logger.Config{
		LogLevel: logger.Error,
	})
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Db.User, cfg.Db.Passwd, cfg.Db.Host, cfg.Db.Name),
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: dblogger.LogMode(logger.Silent),
	})

	if err != nil {
		l.Error("Database connection failed")
		panic(err)
	}
	return &mysqlDatabase{db: db}
}

func (p *mysqlDatabase) GetDb() *gorm.DB {
	return p.db
}
