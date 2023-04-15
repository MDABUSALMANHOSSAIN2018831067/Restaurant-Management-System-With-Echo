package config

import (
	"log"

	"github.com/spf13/viper"
)

var LocalConfig *Config

type Config struct {
	DBUser    string `mapstructure:"DBUSER"`
	DBPass    string `mapstructure:"DBPASS"`
	DbName    string `mapstructure:"DBNAME"`
	Port      string `mapstructure:"PORT"`
	SECRETKEY string `mapstructure:"SECRET_KEY"`
}

func InitConfig() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	var config *Config

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error reading env file", err)
	}

	return config
}

func SetConfig() {
	LocalConfig = InitConfig()
}
