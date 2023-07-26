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

type AddressHandlerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type addressHandler struct {
	application application.AddressService
}

func NewAddressHandler(application application.AddressService) AddressHandlerInterface {
	return &addressHandler{application}
}

func (addressHandler *addressHandler) Create(c *gin.Context) {
	var input request.AddressCreateRequest
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	address, err := addressHandler.application.Create(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save address !", http.StatusOK, helper.SUCCESS, response.FormatAddress(address))
	c.JSON(http.StatusOK, response)
	return
}

func (addressHandler *addressHandler) Update(c *gin.Context) {
	var inputID request.AddressFindById
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	var inputData request.AddressUpdateRequest
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	updatedAddress, err := addressHandler.application.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success update address", http.StatusOK, helper.SUCCESS, response.FormatAddress(updatedAddress))
	c.JSON(http.StatusOK, response)
}

func (addressHandler *addressHandler) FindById(c *gin.Context) {
	var input request.AddressFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	address, err := addressHandler.application.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("Success get address !", http.StatusOK, helper.SUCCESS, response.FormatAddress(address))
	c.JSON(http.StatusOK, response)
}

func (addressHandler *addressHandler) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	addresss, err := addressHandler.application.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of addresss", http.StatusOK, helper.SUCCESS, response.FormatAddresss(addresss))
	c.JSON(http.StatusOK, response)
}

func (addressHandler *addressHandler) Delete(c *gin.Context) {
	var input request.AddressFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete address", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	address, err := addressHandler.application.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete address", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete address !", http.StatusOK, helper.SUCCESS, response.FormatAddress(address))
	c.JSON(http.StatusOK, response)
}
