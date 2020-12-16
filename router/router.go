package router

import (
	"github.com/daystram/go-gin-gorm-boilerplate/controllers/middleware"
	"github.com/daystram/go-gin-gorm-boilerplate/controllers/v1"

	"github.com/gin-gonic/gin"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	v1route := router.Group("/v1")
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
			user.GET("/:id", v1.GETUser)
			user.POST("/", v1.POSTUser)
		}
	}
	return:Q
}
