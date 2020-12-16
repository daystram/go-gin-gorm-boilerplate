package config

import (
	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	Port        int
	Environment string
	Debug       bool

	Domain    string
	JWTSecret string
}

func InitializeAppConfig() {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Port
	AppConfig.Port = viper.GetInt("port")

	// Environment
	if AppConfig.Environment = viper.GetString("environment"); AppConfig.Environment == "" {
		panic("environment is missing in config.yaml")
	}

	// Domain
	if AppConfig.Domain = viper.GetString("domain"); AppConfig.Domain == "" {
		panic("domain is missing in config.yaml")
	}

	// Debug
	AppConfig.Debug = viper.GetBool("debug")

	// JWTSecret
	if AppConfig.JWTSecret = viper.GetString("secret"); AppConfig.JWTSecret == "" {
		panic("secret is missing in config.yaml")
	}
}
