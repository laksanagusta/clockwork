package handler

import (
	"clockwork-server/application"
	"clockwork-server/domain/model"
	"clockwork-server/helper"
	"clockwork-server/interfaces/api/request"
	"clockwork-server/interfaces/api/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const SAVE_PRODUCT_FAILED = "Failed to save product"
const GET_PRODUCT_FAILED = "Failed to get product"

type ProductHandlerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindByCode(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type productHandler struct {
	application application.ProductService
}

func NewProductHandler(application application.ProductService) ProductHandlerInterface {
	return &productHandler{application}
}

func (_productHandler *productHandler) Create(c *gin.Context) {
	var input request.ProductCreateInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	currentUser := c.MustGet("currentUser").(model.User)
	input.User = currentUser

	product, err := _productHandler.application.Create(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save product !", http.StatusOK, helper.SUCCESS, response.FormatProduct(product))
	c.JSON(http.StatusOK, response)
}

func (_productHandler *productHandler) Update(c *gin.Context) {
	var inputID request.ProductFindById
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	var inputData request.ProductUpdateInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	updatedProduct, err := _productHandler.application.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create product", http.StatusOK, helper.SUCCESS, response.FormatProduct(updatedProduct))
	c.JSON(http.StatusOK, response)
}

func (_productHandler *productHandler) FindById(c *gin.Context) {
	var input request.ProductFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	product, err := _productHandler.application.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
		return
	}

	response := helper.APIResponse("Success get product !", http.StatusOK, helper.SUCCESS, response.FormatProduct(product))
	c.JSON(http.StatusOK, response)
}

func (_productHandler *productHandler) FindByCode(c *gin.Context) {
	var input request.ProductFindBySerialNumber
	err := c.ShouldBindUri(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	productSingle, err := _productHandler.application.FindBySerialNumber(input.SerialNumber)
	if err != nil {
		helper.NotFoundResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
		return
	}

	response := helper.APIResponse("Success get product !", http.StatusOK, helper.SUCCESS, response.FormatProduct(productSingle))
	c.JSON(http.StatusOK, response)
}

func (_productHandler *productHandler) FindAll(c *gin.Context) {
	urlQuery := c.Request.URL.Query()
	page, _ := strconv.Atoi(urlQuery.Get("page"))
	limit, _ := strconv.Atoi(urlQuery.Get("limit"))
	title := urlQuery.Get("title")
	categoryId := urlQuery.Get("category_id")

	products, err := _productHandler.application.FindAll(page, limit, title, categoryId)
	if err != nil {
		helper.NotFoundResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
		return
	}

	response := helper.APIResponse("List of products", http.StatusOK, helper.SUCCESS, response.FormatProducts(products))
	c.JSON(http.StatusOK, response)
}

func (_productHandler *productHandler) Delete(c *gin.Context) {
	var input request.ProductFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	product, err := _productHandler.application.Delete(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.DELETE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success delete product !", http.StatusOK, helper.SUCCESS, response.FormatProduct(product))
	c.JSON(http.StatusOK, response)
}
