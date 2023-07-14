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

type RackControllerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type rackController struct {
	service service.RackService
}

func NewRackController(service service.RackService) RackControllerInterface {
	return &rackController{service}
}

func (rackController *rackController) Create(c *gin.Context) {
	var input request.RackCreateInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	rack, err := rackController.service.Create(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save rack !", http.StatusOK, helper.SUCCESS, response.FormatRack(rack))
	c.JSON(http.StatusOK, response)
	return
}

func (rackController *rackController) Update(c *gin.Context) {
	var inputID request.RackFindById
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	var inputData request.RackUpdateInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	updatedRack, err := rackController.service.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create rack", http.StatusOK, helper.SUCCESS, response.FormatRack(updatedRack))
	c.JSON(http.StatusOK, response)
}

func (rackController *rackController) FindById(c *gin.Context) {
	var input request.RackFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	rack, err := rackController.service.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("Success get rack !", http.StatusOK, helper.SUCCESS, response.FormatRack(rack))
	c.JSON(http.StatusOK, response)
}

func (rackController *rackController) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	racks, err := rackController.service.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of racks", http.StatusOK, helper.SUCCESS, response.FormatRacks(racks))
	c.JSON(http.StatusOK, response)
}

func (rackController *rackController) Delete(c *gin.Context) {
	var input request.RackFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete rack", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	rack, err := rackController.service.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete rack", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete rack !", http.StatusOK, helper.SUCCESS, response.FormatRack(rack))
	c.JSON(http.StatusOK, response)
}
