package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorValidation(err error, c *gin.Context, message string) {
	errors := FormatValidationError(err)
	errorMessageData := gin.H{"errors": errors}
	response := APIResponse(message, http.StatusUnprocessableEntity, "error", errorMessageData)
	c.JSON(http.StatusUnprocessableEntity, response)
}

func ErrorResponse(err error, c *gin.Context, message string) {
	errorMessageData := gin.H{"errors": err.Error()}
	response := APIResponse(message, http.StatusUnprocessableEntity, "error", errorMessageData)
	c.JSON(http.StatusUnprocessableEntity, response)
}

func NotFoundResponse(err error, c *gin.Context, message string) {
	errors := FormatValidationError(err)
	errorMessageData := gin.H{"errors": errors}
	response := APIResponse(message, http.StatusUnprocessableEntity, "error", errorMessageData)
	c.JSON(http.StatusUnprocessableEntity, response)
}

func UnauthorizedResponse(err error, c *gin.Context, message string) {
	errors := FormatValidationError(err)
	errorMessageData := gin.H{"errors": errors}
	response := APIResponse(message, http.StatusUnprocessableEntity, "error", errorMessageData)
	c.JSON(http.StatusUnprocessableEntity, response)
}

func SuccessResponse(err error, c *gin.Context, message string) {
	errors := FormatValidationError(err)
	errorMessageData := gin.H{"errors": errors}
	response := APIResponse(message, http.StatusUnprocessableEntity, "error", errorMessageData)
	c.JSON(http.StatusUnprocessableEntity, response)
}
