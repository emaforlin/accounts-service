package config

import (
	"flag"

	"github.com/spf13/viper"
)

var migrate bool

type Config struct {
	App App
	Db  Db
}

type App struct {
	ApiVersion string
	Port       uint16
}

type Db struct {
	Migrate bool
	Uri     string
	Name    string
	User    string
	Passwd  string
	Host    string
}

func LoadConfig() *Config {
	flag.BoolVar(&migrate, "migrate", false, "enables database auto migration")
	flag.Parse()

	return &Config{
		App: App{
			ApiVersion: viper.GetString("service.api"),
			Port:       50014,
		},
		Db: Db{
			Migrate: migrate,
			Name:    viper.GetString("database.name"),
			User:    viper.GetString("database.user"),
			Passwd:  viper.GetString("database.password"),
			Host:    viper.GetString("database.host"),
		},
	}
}

func InitViper(filename string) {
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		panic("error reading config file")
	}
}
