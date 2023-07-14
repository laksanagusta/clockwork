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

type OrganizationControllerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type organizationController struct {
	service service.OrganizationService
}

func NewOrganizationController(service service.OrganizationService) OrganizationControllerInterface {
	return &organizationController{service}
}

func (organizationController *organizationController) Create(c *gin.Context) {
	var input request.OrganizationCreateInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	organization, err := organizationController.service.Create(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save organization !", http.StatusOK, helper.SUCCESS, response.FormatOrganization(organization))
	c.JSON(http.StatusOK, response)
	return
}

func (organizationController *organizationController) Update(c *gin.Context) {
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

	updatedOrganization, err := organizationController.service.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create organization", http.StatusOK, helper.SUCCESS, response.FormatOrganization(updatedOrganization))
	c.JSON(http.StatusOK, response)
}

func (organizationController *organizationController) FindById(c *gin.Context) {
	var input request.OrganizationFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	organization, err := organizationController.service.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("Success get organization !", http.StatusOK, helper.SUCCESS, response.FormatOrganization(organization))
	c.JSON(http.StatusOK, response)
}

func (organizationController *organizationController) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	organizations, err := organizationController.service.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of organizations", http.StatusOK, helper.SUCCESS, response.FormatOrganizations(organizations))
	c.JSON(http.StatusOK, response)
}

func (organizationController *organizationController) Delete(c *gin.Context) {
	var input request.OrganizationFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete organization", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	organization, err := organizationController.service.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete organization", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete organization !", http.StatusOK, helper.SUCCESS, response.FormatOrganization(organization))
	c.JSON(http.StatusOK, response)
}
