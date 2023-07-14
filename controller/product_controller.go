package controller

import (
	"clockwork-server/helper"
	"clockwork-server/model"
	"clockwork-server/request"
	"clockwork-server/response"
	"clockwork-server/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const SAVE_PRODUCT_FAILED = "Failed to save product"
const GET_PRODUCT_FAILED = "Failed to get product"

type ProductControllerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindByCode(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type productController struct {
	service service.ProductService
}

func NewProductController(service service.ProductService) ProductControllerInterface {
	return &productController{service}
}

func (_productController *productController) Create(c *gin.Context) {
	var input request.ProductCreateInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	currentUser := c.MustGet("currentUser").(model.User)
	input.User = currentUser

	product, err := _productController.service.Create(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save product !", http.StatusOK, helper.SUCCESS, response.FormatProduct(product))
	c.JSON(http.StatusOK, response)
}

func (_productController *productController) Update(c *gin.Context) {
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

	updatedProduct, err := _productController.service.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create product", http.StatusOK, helper.SUCCESS, response.FormatProduct(updatedProduct))
	c.JSON(http.StatusOK, response)
}

func (_productController *productController) FindById(c *gin.Context) {
	var input request.ProductFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	product, err := _productController.service.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
		return
	}

	response := helper.APIResponse("Success get product !", http.StatusOK, helper.SUCCESS, response.FormatProduct(product))
	c.JSON(http.StatusOK, response)
}

func (_productController *productController) FindByCode(c *gin.Context) {
	var input request.ProductFindBySerialNumber
	err := c.ShouldBindUri(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	productSingle, err := _productController.service.FindBySerialNumber(input.SerialNumber)
	if err != nil {
		helper.NotFoundResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
		return
	}

	response := helper.APIResponse("Success get product !", http.StatusOK, helper.SUCCESS, response.FormatProduct(productSingle))
	c.JSON(http.StatusOK, response)
}

func (_productController *productController) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	s := q.Get("q")
	products, err := _productController.service.FindAll(page, pageSize, s)
	if err != nil {
		helper.NotFoundResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
		return
	}

	response := helper.APIResponse("List of products", http.StatusOK, helper.SUCCESS, response.FormatProducts(products))
	c.JSON(http.StatusOK, response)
}

func (_productController *productController) Delete(c *gin.Context) {
	var input request.ProductFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	product, err := _productController.service.Delete(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.DELETE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success delete product !", http.StatusOK, helper.SUCCESS, response.FormatProduct(product))
	c.JSON(http.StatusOK, response)
}
