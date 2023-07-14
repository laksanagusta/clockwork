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

type LocationControllerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type locationController struct {
	service service.LocationService
}

func NewLocationController(service service.LocationService) LocationControllerInterface {
	return &locationController{service}
}

func (locationController *locationController) Create(c *gin.Context) {
	var input request.LocationCreateInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	location, err := locationController.service.Create(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save location !", http.StatusOK, helper.SUCCESS, response.FormatLocation(location))
	c.JSON(http.StatusOK, response)
	return
}

func (locationController *locationController) Update(c *gin.Context) {
	var inputID request.LocationFindById
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	var inputData request.LocationUpdateInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	updatedLocation, err := locationController.service.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create location", http.StatusOK, helper.SUCCESS, response.FormatLocation(updatedLocation))
	c.JSON(http.StatusOK, response)
}

func (locationController *locationController) FindById(c *gin.Context) {
	var input request.LocationFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	location, err := locationController.service.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("Success get location !", http.StatusOK, helper.SUCCESS, response.FormatLocation(location))
	c.JSON(http.StatusOK, response)
}

func (locationController *locationController) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	locations, err := locationController.service.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of locations", http.StatusOK, helper.SUCCESS, response.FormatLocations(locations))
	c.JSON(http.StatusOK, response)
}

func (locationController *locationController) Delete(c *gin.Context) {
	var input request.LocationFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete location", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	location, err := locationController.service.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete location", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete location !", http.StatusOK, helper.SUCCESS, response.FormatLocation(location))
	c.JSON(http.StatusOK, response)
}
