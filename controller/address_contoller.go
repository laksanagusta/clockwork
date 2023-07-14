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

type AddressControllerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type addressController struct {
	service service.AddressService
}

func NewAddressController(service service.AddressService) AddressControllerInterface {
	return &addressController{service}
}

func (addressController *addressController) Create(c *gin.Context) {
	var input request.AddressCreateRequest
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	address, err := addressController.service.Create(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save address !", http.StatusOK, helper.SUCCESS, response.FormatAddress(address))
	c.JSON(http.StatusOK, response)
	return
}

func (addressController *addressController) Update(c *gin.Context) {
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

	updatedAddress, err := addressController.service.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success update address", http.StatusOK, helper.SUCCESS, response.FormatAddress(updatedAddress))
	c.JSON(http.StatusOK, response)
}

func (addressController *addressController) FindById(c *gin.Context) {
	var input request.AddressFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	address, err := addressController.service.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("Success get address !", http.StatusOK, helper.SUCCESS, response.FormatAddress(address))
	c.JSON(http.StatusOK, response)
}

func (addressController *addressController) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	addresss, err := addressController.service.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of addresss", http.StatusOK, helper.SUCCESS, response.FormatAddresss(addresss))
	c.JSON(http.StatusOK, response)
}

func (addressController *addressController) Delete(c *gin.Context) {
	var input request.AddressFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete address", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	address, err := addressController.service.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete address", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete address !", http.StatusOK, helper.SUCCESS, response.FormatAddress(address))
	c.JSON(http.StatusOK, response)
}
