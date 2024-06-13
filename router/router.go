package router

import (
	"github.com/broboredo/locapp-api/config"
	docs "github.com/broboredo/locapp-api/docs"
	"github.com/broboredo/locapp-api/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() {
	config.LoadEnv()
	configSwagger()

	router := gin.Default()
	initRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	protected := router.Group("/static/images/products")
	protected.Use(middleware.SecurityToken())
	{
		protected.Static("/", "./static/images/products")
	}

	router.Run(":8080")
}

func configSwagger() {
	docs.SwaggerInfo.Title = "LocApp API Doc"
	docs.SwaggerInfo.Description = "LocApp is a project for you manage your rentable products"
	docs.SwaggerInfo.Version = "0.1.0"
	docs.SwaggerInfo.BasePath = "api/v1"
}
