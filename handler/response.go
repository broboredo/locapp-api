package handler

import (
	"net/http"

	"github.com/broboredo/locapp-api/schemas"
	"github.com/gin-gonic/gin"
)

func SendError(c *gin.Context, statusCode int, message string) {
	c.Header("Content-type", "application/json")
	c.JSON(statusCode, gin.H{
		"error": message,
	})
}

func SendSuccess(c *gin.Context, data interface{}, statusCode ...int) {
	c.Header("Content-type", "application/json")

	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	c.JSON(code, gin.H{
		"data": data,
	})
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type ProductResponse struct {
	Data schemas.ProductResponse `json:"data"`
}

type ListProductResponse struct {
	Data []schemas.ProductResponse `json:"data"`
}

type PaginationResponse struct {
	Data       interface{} `json:"data"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	TotalPages int         `json:"totalPages"`
}
