package main

import (
	"fmt"

	"github.com/daystram/go-gin-gorm-boilerplate/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitializeAppConfig()
	if !config.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Run(fmt.Sprintf(":%d", config.AppConfig.Port))
}
