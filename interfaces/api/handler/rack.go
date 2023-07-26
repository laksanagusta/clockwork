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

type RackHandlerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type rackHandler struct {
	application application.RackService
}

func NewRackHandler(application application.RackService) RackHandlerInterface {
	return &rackHandler{application}
}

func (rackHandler *rackHandler) Create(c *gin.Context) {
	var input request.RackCreateInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	rack, err := rackHandler.application.Create(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save rack !", http.StatusOK, helper.SUCCESS, response.FormatRack(rack))
	c.JSON(http.StatusOK, response)
	return
}

func (rackHandler *rackHandler) Update(c *gin.Context) {
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

	updatedRack, err := rackHandler.application.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create rack", http.StatusOK, helper.SUCCESS, response.FormatRack(updatedRack))
	c.JSON(http.StatusOK, response)
}

func (rackHandler *rackHandler) FindById(c *gin.Context) {
	var input request.RackFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	rack, err := rackHandler.application.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("Success get rack !", http.StatusOK, helper.SUCCESS, response.FormatRack(rack))
	c.JSON(http.StatusOK, response)
}

func (rackHandler *rackHandler) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	racks, err := rackHandler.application.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of racks", http.StatusOK, helper.SUCCESS, response.FormatRacks(racks))
	c.JSON(http.StatusOK, response)
}

func (rackHandler *rackHandler) Delete(c *gin.Context) {
	var input request.RackFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete rack", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	rack, err := rackHandler.application.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete rack", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete rack !", http.StatusOK, helper.SUCCESS, response.FormatRack(rack))
	c.JSON(http.StatusOK, response)
}
