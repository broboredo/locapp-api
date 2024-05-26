package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ValidateProductRequest(req interface{}, context *gin.Context) error {
	validate := validator.New()

	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}

	if req == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Request body is empty"})
		return nil
	}

	if err := validate.Struct(req); err != nil {
		Logger.Errorf("error validating CreateProductRequest: %v", err)
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			Logger.Errorf("error validating CreateProductRequest internal: %v", err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return err
		}

		errs := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Field %s is %s", err.Field(), err.Tag())

			if err.Param() != "" {
				errorMessage += fmt.Sprintf(": %s", err.Param())
			}

			errs[err.Field()] = errorMessage
		}

		Logger.Errorf("error validating CreateProductRequest response: %v", err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": errs})

		return err
	}

	return nil
}
