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

type LocationHandlerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type locationHandler struct {
	application application.LocationService
}

func NewLocationHandler(application application.LocationService) LocationHandlerInterface {
	return &locationHandler{application}
}

func (locationHandler *locationHandler) Create(c *gin.Context) {
	var input request.LocationCreateInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	location, err := locationHandler.application.Create(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save location !", http.StatusOK, helper.SUCCESS, response.FormatLocation(location))
	c.JSON(http.StatusOK, response)
	return
}

func (locationHandler *locationHandler) Update(c *gin.Context) {
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

	updatedLocation, err := locationHandler.application.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create location", http.StatusOK, helper.SUCCESS, response.FormatLocation(updatedLocation))
	c.JSON(http.StatusOK, response)
}

func (locationHandler *locationHandler) FindById(c *gin.Context) {
	var input request.LocationFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	location, err := locationHandler.application.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("Success get location !", http.StatusOK, helper.SUCCESS, response.FormatLocation(location))
	c.JSON(http.StatusOK, response)
}

func (locationHandler *locationHandler) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	locations, err := locationHandler.application.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of locations", http.StatusOK, helper.SUCCESS, response.FormatLocations(locations))
	c.JSON(http.StatusOK, response)
}

func (locationHandler *locationHandler) Delete(c *gin.Context) {
	var input request.LocationFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete location", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	location, err := locationHandler.application.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete location", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete location !", http.StatusOK, helper.SUCCESS, response.FormatLocation(location))
	c.JSON(http.StatusOK, response)
}
