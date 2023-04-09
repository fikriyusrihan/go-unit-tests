package config

import (
	"github.com/spf13/viper"
	"log"
)

type config struct {
	Server struct {
		Port string
	}

	Database struct {
		Username string
		Password string
		Host     string
		Port     string
		DBName   string
		SSLMode  string
	}

	JWT struct {
		SignatureKey string
	}
}

var C config

func ReadConfig() {
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(Config); err != nil {
		log.Fatalln(err)
	}
}
