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

type CartControllerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
	CheckActiveCart(c *gin.Context)
}

type cartController struct {
	service service.CartService
}

func NewCartController(service service.CartService) CartControllerInterface {
	return &cartController{service}
}

func (cartController *cartController) Create(c *gin.Context) {
	customerId := c.MustGet("curentCustomer").(model.Customer).ID

	cart, err := cartController.service.Create(int(customerId))
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save cart !", http.StatusOK, helper.SUCCESS, response.FormatCart(cart))
	c.JSON(http.StatusOK, response)
	return
}

func (cartController *cartController) Update(c *gin.Context) {
	var inputID request.CartFindById

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	var inputData request.CartUpdateRequest
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	updatedCart, err := cartController.service.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create order item", http.StatusOK, helper.SUCCESS, response.FormatCart(updatedCart))
	c.JSON(http.StatusOK, response)
}

func (cartController *cartController) FindById(c *gin.Context) {
	var input request.CartFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	cart, err := cartController.service.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("Success get order item !", http.StatusOK, helper.SUCCESS, response.FormatCart(cart))
	c.JSON(http.StatusOK, response)
}

func (cartController *cartController) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	carts, err := cartController.service.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of order items", http.StatusOK, helper.SUCCESS, response.FormatCarts(carts))
	c.JSON(http.StatusOK, response)
}

func (cartController *cartController) Delete(c *gin.Context) {
	var input request.CartFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete cart", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	cart, err := cartController.service.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete cart", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete cart !", http.StatusOK, helper.SUCCESS, response.FormatCart(cart))
	c.JSON(http.StatusOK, response)
}

func (cc *cartController) CheckActiveCart(c *gin.Context) {
	customerId := c.MustGet("currentCustomer").(int)
	cart, err := cc.service.CheckActiveCart(customerId)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
		return
	}

	response := helper.APIResponse("Success get active cart !", http.StatusOK, helper.SUCCESS, response.FormatCart(cart))
	c.JSON(http.StatusOK, response)
}
