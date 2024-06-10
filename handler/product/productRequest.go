package product

import (
	"errors"
	"github.com/broboredo/locapp-api/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=3,max=100"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"required,numeric,gt=0"`
	Quantity    int     `json:"quantity" validate:"required,numeric,gt=0"`
}

type UpdateProductRequest struct {
	Name        string   `json:"name" validate:"min=3,max=100"`
	Description string   `json:"description"`
	Price       *float32 `json:"price" validate:"numeric"`
	Quantity    *int     `json:"quantity" validate:"numeric"`
	ImagePath   string   `json:"image"`
}

type IdIsRequired struct{}

func (req *CreateProductRequest) Validate(context *gin.Context) error {
	return handler.ValidateProductRequest(req, context)
}

func (req *UpdateProductRequest) Validate(context *gin.Context) error {
	r := IdIsRequired{}
	err := r.Validate(context)
	if err != nil {
		return err
	}

	return handler.ValidateProductRequest(req, context)
}

func (req *IdIsRequired) Validate(context *gin.Context) error {
	id := context.Param("id")
	if id == "" {
		handler.SendError(context, http.StatusBadRequest, "no product id provided")
		return errors.New("no product id provided")
	}

	return nil
}
