package handlers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/daystram/go-gin-gorm-boilerplate/config"
	"github.com/daystram/go-gin-gorm-boilerplate/datatransfers"
	"github.com/daystram/go-gin-gorm-boilerplate/models"
)

var Handler HandlerFunc

type HandlerFunc interface {
	AuthenticateUser(credentials datatransfers.UserLogin) (token string, err error)
	RegisterUser(credentials datatransfers.UserSignup) (err error)

	RetrieveUser(username string) (user models.User, err error)
	UpdateUser(id uint, user datatransfers.UserUpdate) (err error)
}

type module struct {
	db *dbEntity
}

type dbEntity struct {
	conn      *gorm.DB
	userOrmer models.UserOrmer
}

func InitializeHandler() (err error) {
	// Initialize DB
	var db *gorm.DB
	switch config.AppConfig.DBDRIVER {
	case "postgres":
		db, err = gorm.Open(postgres.Open(
			fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
				config.AppConfig.DBHost, config.AppConfig.DBPort, config.AppConfig.DBDatabase,
				config.AppConfig.DBUsername, config.AppConfig.DBPassword),
		), &gorm.Config{})
		if err != nil {
			log.Println("[INIT] failed connecting to PostgreSQL")
			return
		}
		log.Println("[INIT] connected to PostgreSQL")
	case "mysql":
		db, err = gorm.Open(mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				config.AppConfig.DBUsername, config.AppConfig.DBPassword, config.AppConfig.DBHost,
				config.AppConfig.DBPort, config.AppConfig.DBDatabase),
		), &gorm.Config{})
		if err != nil {
			log.Println("[INIT] failed connecting to MySQL")
			return
		}
		log.Println("[INIT] connected to MySQL")
	default:
		log.Println("DB Driver not specified or not supported. Possible options are: 'postgres' or 'mysql'")
		os.Exit(1)
	}

	// Compose handler modules
	Handler = &module{
		db: &dbEntity{
			conn:      db,
			userOrmer: models.NewUserOrmer(db),
		},
	}
	return
}
