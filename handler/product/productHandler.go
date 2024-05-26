package product

import (
	"fmt"
	"github.com/broboredo/locapp-api/handler"
	"github.com/broboredo/locapp-api/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary Create Product
// @Description Create a new product
// @Tags Product
// @Accept json
// @Produce json
// @Param request body CreateProductRequest true "Request body"
// @Success 200 {object} handler.ProductResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 404 {object} handler.ErrorResponse
// @Failure 422 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /products [post]
func Create(context *gin.Context) {
	var req CreateProductRequest
	if err := req.Validate(context); err != nil {
		return
	}

	product := schemas.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
	}

	if err := handler.Db.Create(&product).Error; err != nil {
		handler.Logger.Errorf("error when try create new product %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	handler.SendSuccess(context, product, http.StatusCreated)
}

// @Summary Update Product
// @Description Update a product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param request body UpdateProductRequest true "Request body"
// @Success 200 {object} handler.ProductResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 404 {object} handler.ErrorResponse
// @Failure 422 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /products/:id [put]
func Update(context *gin.Context) {
	var req UpdateProductRequest
	if err := req.Validate(context); err != nil {
		return
	}

	product := schemas.Product{}

	id := context.Param("id")
	if err := handler.Db.First(&product, id).Error; err != nil {
		handler.SendError(context, http.StatusNotFound, fmt.Sprintf("Product with id: %s not found", id))
		return
	}

	if req.Name != "" {
		product.Name = req.Name
	}

	if req.Price != nil {
		product.Price = *req.Price
	}

	if req.Description != "" {
		product.Description = req.Description
	}

	if req.Quantity != nil {
		product.Quantity = *req.Quantity
	}

	// Save opening
	if err := handler.Db.Save(&product).Error; err != nil {
		handler.Logger.Errorf("error updating opening: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, "error updating opening")
		return
	}

	handler.SendSuccess(context, product)
}

// @Summary List Product
// @Description List products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} handler.ListProductResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 404 {object} handler.ErrorResponse
// @Failure 422 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /products [get]
func List(context *gin.Context) {
	var products []schemas.Product

	if err := handler.Db.Find(&products).Error; err != nil {
		handler.SendError(context, http.StatusInternalServerError, "error listing products")
		return
	}

	handler.SendSuccess(context, products)
}

// @Summary Find Product
// @Description Find a product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} handler.ProductResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 404 {object} handler.ErrorResponse
// @Failure 422 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /products/:id [get]
func Find(context *gin.Context) {
	var req IdIsRequired
	if err := req.Validate(context); err != nil {
		return
	}

	product := schemas.Product{}

	id := context.Param("id")
	if err := handler.Db.First(&product, id).Error; err != nil {
		handler.SendError(context, http.StatusNotFound, fmt.Sprintf("Product with id: %s not found", id))
		return
	}

	handler.SendSuccess(context, product)
}

// @Summary Delete Product
// @Description Delete a product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} handler.ProductResponse
// @Failure 400 {object} handler.ErrorResponse
// @Failure 404 {object} handler.ErrorResponse
// @Failure 500 {object} handler.ErrorResponse
// @Router /products/:id [delete]
func Delete(context *gin.Context) {
	var req IdIsRequired
	if err := req.Validate(context); err != nil {
		return
	}

	opening := schemas.Product{}

	db := handler.Db
	id := context.Param("id")
	if err := db.First(&opening, id).Error; err != nil {
		handler.SendError(context, http.StatusNotFound, fmt.Sprintf("Product with id: %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		handler.SendError(context, http.StatusInternalServerError, fmt.Sprintf("error deleting Product with id: %s", id))
		return
	}

	handler.SendSuccess(context, opening)
}
