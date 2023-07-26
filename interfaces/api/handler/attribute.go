package handler

import (
	"clockwork-server/application"
	"clockwork-server/helper"
	"clockwork-server/interfaces/api/request"
	"clockwork-server/interfaces/api/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AttributeHandlerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type attributeHandler struct {
	application application.AttributeService
}

func NewAttributeHandler(application application.AttributeService) AttributeHandlerInterface {
	return &attributeHandler{application}
}

func (attributeHandler *attributeHandler) Create(c *gin.Context) {
	var input request.AttributeCreateRequest
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	attribute, err := attributeHandler.application.Create(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save attribute !", http.StatusOK, helper.SUCCESS, response.FormatAttribute(attribute))
	c.JSON(http.StatusOK, response)
	return
}

func (attributeHandler *attributeHandler) Update(c *gin.Context) {
	var inputID request.AttributeFindById
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	var inputData request.AttributeUpdateRequest
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	updatedAttribute, err := attributeHandler.application.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create attribute", http.StatusOK, helper.SUCCESS, response.FormatAttribute(updatedAttribute))
	c.JSON(http.StatusOK, response)
}

func (attributeHandler *attributeHandler) FindById(c *gin.Context) {
	var input request.AttributeFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	attribute, err := attributeHandler.application.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("Success get attribute !", http.StatusOK, helper.SUCCESS, response.FormatAttribute(attribute))
	c.JSON(http.StatusOK, response)
}

func (attributeHandler *attributeHandler) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	attributes, err := attributeHandler.application.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of attributes", http.StatusOK, helper.SUCCESS, response.FormatAttributes(attributes))
	c.JSON(http.StatusOK, response)
}

func (attributeHandler *attributeHandler) Delete(c *gin.Context) {
	var input request.AttributeFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete attribute", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	attribute, err := attributeHandler.application.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete attribute", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete attribute !", http.StatusOK, helper.SUCCESS, response.FormatAttribute(attribute))
	c.JSON(http.StatusOK, response)
}
