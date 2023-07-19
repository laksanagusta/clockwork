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

type AttributeItemControllerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type attributeItemController struct {
	service service.AttributeItemService
}

func NewAttributeItemController(service service.AttributeItemService) AttributeItemControllerInterface {
	return &attributeItemController{service}
}

func (attributeItemController *attributeItemController) Create(c *gin.Context) {
	var input request.AttributeItemCreateRequest
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	attributeItem, err := attributeItemController.service.Create(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save attributeItem !", http.StatusOK, helper.SUCCESS, response.FormatAttributeItem(attributeItem))
	c.JSON(http.StatusOK, response)
	return
}

func (attributeItemController *attributeItemController) Update(c *gin.Context) {
	var inputID request.AttributeItemFindById
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	var inputData request.AttributeItemUpdateRequest
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	updatedAttributeItem, err := attributeItemController.service.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create attributeItem", http.StatusOK, helper.SUCCESS, response.FormatAttributeItem(updatedAttributeItem))
	c.JSON(http.StatusOK, response)
}

func (attributeItemController *attributeItemController) FindById(c *gin.Context) {
	var input request.AttributeItemFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	attributeItem, err := attributeItemController.service.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("Success get attributeItem !", http.StatusOK, helper.SUCCESS, response.FormatAttributeItem(attributeItem))
	c.JSON(http.StatusOK, response)
}

func (attributeItemController *attributeItemController) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	attributeItems, err := attributeItemController.service.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of attributeItems", http.StatusOK, helper.SUCCESS, response.FormatAttributeItems(attributeItems))
	c.JSON(http.StatusOK, response)
}

func (attributeItemController *attributeItemController) Delete(c *gin.Context) {
	var input request.AttributeItemFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete attributeItem", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	attributeItem, err := attributeItemController.service.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete attributeItem", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete attributeItem !", http.StatusOK, helper.SUCCESS, response.FormatAttributeItem(attributeItem))
	c.JSON(http.StatusOK, response)
}
