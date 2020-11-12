package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB database `json:"db"`
}

type database struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"database"`
	Port     int    `json:"port"`
}

var Cfg Config

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	postgresConfig := &database{}

	err = viper.Unmarshal(&postgresConfig)
	if err != nil {
		return err
	}

	viper.Set("db", postgresConfig)

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		return err
	}
	return nil
}
