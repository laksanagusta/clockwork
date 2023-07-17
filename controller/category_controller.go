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

type CategoryControllerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type categoryController struct {
	service service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryControllerInterface {
	return &categoryController{service}
}

func (categoryController *categoryController) Create(c *gin.Context) {
	var input request.CategoryCreateInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	category, err := categoryController.service.Create(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save category !", http.StatusOK, helper.SUCCESS, response.FormatCategory(category))
	c.JSON(http.StatusOK, response)
	return
}

func (categoryController *categoryController) Update(c *gin.Context) {
	var inputID request.CategoryFindById
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	var inputData request.CategoryUpdateInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	updatedCategory, err := categoryController.service.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success update category", http.StatusOK, helper.SUCCESS, response.FormatCategory(updatedCategory))
	c.JSON(http.StatusOK, response)
}

func (categoryController *categoryController) FindById(c *gin.Context) {
	var input request.CategoryFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	category, err := categoryController.service.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
		return
	}

	response := helper.APIResponse("Success get category !", http.StatusOK, helper.SUCCESS, response.FormatCategory(category))
	c.JSON(http.StatusOK, response)
}

func (categoryController *categoryController) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	categorys, err := categoryController.service.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of categorys", http.StatusOK, helper.SUCCESS, response.FormatCategories(categorys))
	c.JSON(http.StatusOK, response)
}

func (categoryController *categoryController) Delete(c *gin.Context) {
	var input request.CategoryFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete category", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	category, err := categoryController.service.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete category !", http.StatusOK, helper.SUCCESS, response.FormatCategory(category))
	c.JSON(http.StatusOK, response)
}
