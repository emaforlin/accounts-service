package database

import (
	"fmt"

	"github.com/emaforlin/accounts-service/config"
	"github.com/hashicorp/go-hclog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type mysqlDatabase struct {
	db *gorm.DB
}

func NewMySQLDatabase(cfg *config.Config, l hclog.Logger) Database {
	dblogger := logger.New(l.Named("database").StandardLogger(&hclog.StandardLoggerOptions{}), logger.Config{
		LogLevel: logger.Silent,
	})
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Db.User, cfg.Db.Passwd, cfg.Db.Host, cfg.Db.Name),
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: dblogger,
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
