package router

import (
	"github.com/gin-gonic/gin"

	"github.com/daystram/go-gin-gorm-boilerplate/controllers/middleware"
	"github.com/daystram/go-gin-gorm-boilerplate/controllers/v1"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	v1route := router.Group("/api/v1")
	v1route.Use(
		middleware.CORSMiddleware,
		middleware.AuthMiddleware,
	)
	{
		auth := v1route.Group("/auth")
		{
			auth.POST("/login", v1.POSTLogin)
			auth.POST("/signup", v1.POSTRegister)
		}
		user := v1route.Group("/user")
		{
			user.GET("/:username", v1.GETUser)
			user.PUT("", v1.PUTUser)
		}
	}
	return
}
