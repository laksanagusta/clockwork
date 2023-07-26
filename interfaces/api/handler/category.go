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

type CategoryHandlerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type categoryHandler struct {
	application application.CategoryService
}

func NewCategoryHandler(application application.CategoryService) CategoryHandlerInterface {
	return &categoryHandler{application}
}

func (categoryHandler *categoryHandler) Create(c *gin.Context) {
	var input request.CategoryCreateInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	category, err := categoryHandler.application.Create(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save category !", http.StatusOK, helper.SUCCESS, response.FormatCategory(category))
	c.JSON(http.StatusOK, response)
	return
}

func (categoryHandler *categoryHandler) Update(c *gin.Context) {
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

	updatedCategory, err := categoryHandler.application.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success update category", http.StatusOK, helper.SUCCESS, response.FormatCategory(updatedCategory))
	c.JSON(http.StatusOK, response)
}

func (categoryHandler *categoryHandler) FindById(c *gin.Context) {
	var input request.CategoryFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	category, err := categoryHandler.application.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
		return
	}

	response := helper.APIResponse("Success get category !", http.StatusOK, helper.SUCCESS, response.FormatCategory(category))
	c.JSON(http.StatusOK, response)
}

func (categoryHandler *categoryHandler) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	categorys, err := categoryHandler.application.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of categorys", http.StatusOK, helper.SUCCESS, response.FormatCategories(categorys))
	c.JSON(http.StatusOK, response)
}

func (categoryHandler *categoryHandler) Delete(c *gin.Context) {
	var input request.CategoryFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete category", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	category, err := categoryHandler.application.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete category !", http.StatusOK, helper.SUCCESS, response.FormatCategory(category))
	c.JSON(http.StatusOK, response)
}
