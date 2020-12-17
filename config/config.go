package config

import (
	"log"

	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	Port        int
	Environment string
	Debug       bool

	DBHostname string
	DBPort     int
	DBDatabase string
	DBUsername string
	DBPassword string

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
		log.Fatalf("[INIT] Unable to load configuration. %+v\n", err)
	}

	// Port
	AppConfig.Port = viper.GetInt("port")

	// Environment
	if AppConfig.Environment = viper.GetString("environment"); AppConfig.Environment == "" {
		log.Fatalln("[INIT] environment is missing in config.yaml")
	}

	// Debug
	AppConfig.Debug = viper.GetBool("debug")

	// DBHostname
	if AppConfig.DBHostname = viper.GetString("db_hostname"); AppConfig.DBHostname == "" {
		log.Fatalln("[INIT] db_hostname is missing in config.yaml")
	}

	// DBPort
	AppConfig.DBPort = viper.GetInt("db_port")

	// DBDatabase
	if AppConfig.DBDatabase = viper.GetString("db_database"); AppConfig.DBDatabase == "" {
		log.Fatalln("[INIT] db_database is missing in config.yaml")
	}

	// DBUsername
	if AppConfig.DBUsername = viper.GetString("db_username"); AppConfig.DBUsername == "" {
		log.Fatalln("[INIT] db_username is missing in config.yaml")
	}

	// DBPassword
	if AppConfig.DBPassword = viper.GetString("db_password"); AppConfig.DBPassword == "" {
		log.Fatalln("[INIT] db_password is missing in config.yaml")
	}

	// Domain
	if AppConfig.Domain = viper.GetString("domain"); AppConfig.Domain == "" {
		log.Fatalln("[INIT] domain is missing in config.yaml")
	}

	// JWTSecret
	if AppConfig.JWTSecret = viper.GetString("secret"); AppConfig.JWTSecret == "" {
		log.Fatalln("[INIT] secret is missing in config.yaml")
	}

	log.Printf("[INIT] Configuration loaded from %s\n", viper.ConfigFileUsed())
}
