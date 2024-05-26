package router

import (
	docs "github.com/broboredo/locapp-api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() {
	configSwagger()

	router := gin.Default()
	initRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}

func configSwagger() {
	docs.SwaggerInfo.Title = "LocApp API Doc"
	docs.SwaggerInfo.Description = "LocApp is a project for you manage your rentable products"
	docs.SwaggerInfo.Version = "0.1.0"
	docs.SwaggerInfo.BasePath = "api/v1"
}
