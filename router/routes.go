package router

import (
	"github.com/broboredo/locapp-api/handler"
	"github.com/broboredo/locapp-api/handler/customer"
	"github.com/broboredo/locapp-api/handler/product"
	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	handler.InitHandler()
	v1 := router.Group("api/v1")
	{
		v1.GET("/products", product.List)
		v1.POST("/products", product.Create)
		v1.PUT("/products/:id", product.Update)
		v1.GET("/products/:id", product.Find)
		v1.DELETE("/products/:id", product.Delete)

		v1.GET("/customers", customer.List)
		v1.POST("/customers", customer.Create)
		v1.PUT("/customers/:id", customer.Update)
		v1.GET("/customers/:id", customer.Find)
		v1.DELETE("/customers/:id", customer.Delete)
	}
}
