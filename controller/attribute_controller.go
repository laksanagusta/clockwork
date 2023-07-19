package controller

import (
	"clockwork-server/helper"
	"clockwork-server/request"
	"clockwork-server/response"
	"clockwork-server/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AttributeControllerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type attributeController struct {
	service service.AttributeService
}

func NewAttributeController(service service.AttributeService) AttributeControllerInterface {
	return &attributeController{service}
}

func (attributeController *attributeController) Create(c *gin.Context) {
	var input request.AttributeCreateRequest
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	attribute, err := attributeController.service.Create(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save attribute !", http.StatusOK, helper.SUCCESS, response.FormatAttribute(attribute))
	c.JSON(http.StatusOK, response)
	return
}

func (attributeController *attributeController) Update(c *gin.Context) {
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

	updatedAttribute, err := attributeController.service.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create attribute", http.StatusOK, helper.SUCCESS, response.FormatAttribute(updatedAttribute))
	c.JSON(http.StatusOK, response)
}

func (attributeController *attributeController) FindById(c *gin.Context) {
	var input request.AttributeFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	attribute, err := attributeController.service.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("Success get attribute !", http.StatusOK, helper.SUCCESS, response.FormatAttribute(attribute))
	c.JSON(http.StatusOK, response)
}

func (attributeController *attributeController) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	attributes, err := attributeController.service.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of attributes", http.StatusOK, helper.SUCCESS, response.FormatAttributes(attributes))
	c.JSON(http.StatusOK, response)
}

func (attributeController *attributeController) Delete(c *gin.Context) {
	var input request.AttributeFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete attribute", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	attribute, err := attributeController.service.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete attribute", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete attribute !", http.StatusOK, helper.SUCCESS, response.FormatAttribute(attribute))
	c.JSON(http.StatusOK, response)
}
