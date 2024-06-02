package customer

import (
	"fmt"
	"github.com/broboredo/locapp-api/handler"
	"github.com/broboredo/locapp-api/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(context *gin.Context) {
	var req CreateCustomerRequest
	if err := req.Validate(context); err != nil {
		return
	}

	customer := schemas.Customer{
		Name:    req.Name,
		Phone:   req.Phone,
		Notes:   req.Notes,
		Address: req.Address,
	}

	if err := handler.Db.Create(&customer).Error; err != nil {
		handler.Logger.Errorf("error when try create new customer %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	handler.SendSuccess(context, customer, http.StatusCreated)
}

func Update(context *gin.Context) {
	var req UpdateCustomerRequest
	if err := req.Validate(context); err != nil {
		return
	}

	customer := schemas.Customer{}

	id := context.Param("id")
	if err := handler.Db.First(&customer, id).Error; err != nil {
		handler.SendError(context, http.StatusNotFound, fmt.Sprintf("Customer with id: %s not found", id))
		return
	}

	if req.Name != "" {
		customer.Name = req.Name
	}

	if req.Notes != "" {
		customer.Notes = req.Notes
	}

	if req.Address != "" {
		customer.Address = req.Address
	}

	if req.Phone != "" {
		customer.Phone = req.Phone
	}

	if err := handler.Db.Save(&customer).Error; err != nil {
		handler.Logger.Errorf("error updating opening: %v", err.Error())
		handler.SendError(context, http.StatusInternalServerError, "error updating opening")
		return
	}

	handler.SendSuccess(context, customer)
}

func List(context *gin.Context) {
	var customers []schemas.Customer

	if err := handler.Db.Find(&customers).Error; err != nil {
		handler.SendError(context, http.StatusInternalServerError, "error listing customers")
		return
	}

	handler.SendSuccess(context, customers)
}

func Find(context *gin.Context) {
	var req IdIsRequired
	if err := req.Validate(context); err != nil {
		return
	}

	customer := schemas.Customer{}

	id := context.Param("id")
	if err := handler.Db.First(&customer, id).Error; err != nil {
		handler.SendError(context, http.StatusNotFound, fmt.Sprintf("Customer with id: %s not found", id))
		return
	}

	handler.SendSuccess(context, customer)
}

func Delete(context *gin.Context) {
	var req IdIsRequired
	if err := req.Validate(context); err != nil {
		return
	}

	customer := schemas.Customer{}

	db := handler.Db
	id := context.Param("id")
	if err := db.First(&customer, id).Error; err != nil {
		handler.SendError(context, http.StatusNotFound, fmt.Sprintf("Customer with id: %s not found", id))
		return
	}

	if err := db.Delete(&customer).Error; err != nil {
		handler.SendError(context, http.StatusInternalServerError, fmt.Sprintf("error deleting Customer with id: %s", id))
		return
	}

	handler.SendSuccess(context, customer)
}
