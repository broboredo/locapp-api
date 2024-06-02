package customer

import (
	"errors"
	"github.com/broboredo/locapp-api/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateCustomerRequest struct {
	Name    string `json:"name" validate:"required,min=3,max=255"`
	Address string `json:"description"`
	Phone   string `json:"phone"`
	Notes   string `json:"notes"`
}

type UpdateCustomerRequest struct {
	Name    string `json:"name" validate:"min=3,max=255"`
	Address string `json:"description"`
	Phone   string `json:"phone"`
	Notes   string `json:"notes"`
}

type IdIsRequired struct{}

func (req *CreateCustomerRequest) Validate(context *gin.Context) error {
	return handler.ValidateProductRequest(req, context)
}

func (req *UpdateCustomerRequest) Validate(context *gin.Context) error {
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
		handler.SendError(context, http.StatusBadRequest, "no customer id provided")
		return errors.New("no customer id provided")
	}

	return nil
}
