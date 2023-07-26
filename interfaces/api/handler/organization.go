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

type OrganizationHandlerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type organizationHandler struct {
	application application.OrganizationService
}

func NewOrganizationHandler(application application.OrganizationService) OrganizationHandlerInterface {
	return &organizationHandler{application}
}

func (organizationHandler *organizationHandler) Create(c *gin.Context) {
	var input request.OrganizationCreateInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	organization, err := organizationHandler.application.Create(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save organization !", http.StatusOK, helper.SUCCESS, response.FormatOrganization(organization))
	c.JSON(http.StatusOK, response)
	return
}

func (organizationHandler *organizationHandler) Update(c *gin.Context) {
	var inputID request.OrganizationFindById
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	var inputData request.OrganizationUpdateInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	updatedOrganization, err := organizationHandler.application.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create organization", http.StatusOK, helper.SUCCESS, response.FormatOrganization(updatedOrganization))
	c.JSON(http.StatusOK, response)
}

func (organizationHandler *organizationHandler) FindById(c *gin.Context) {
	var input request.OrganizationFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	organization, err := organizationHandler.application.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("Success get organization !", http.StatusOK, helper.SUCCESS, response.FormatOrganization(organization))
	c.JSON(http.StatusOK, response)
}

func (organizationHandler *organizationHandler) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	organizations, err := organizationHandler.application.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of organizations", http.StatusOK, helper.SUCCESS, response.FormatOrganizations(organizations))
	c.JSON(http.StatusOK, response)
}

func (organizationHandler *organizationHandler) Delete(c *gin.Context) {
	var input request.OrganizationFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete organization", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	organization, err := organizationHandler.application.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete organization", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete organization !", http.StatusOK, helper.SUCCESS, response.FormatOrganization(organization))
	c.JSON(http.StatusOK, response)
}
