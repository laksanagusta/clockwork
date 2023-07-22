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

type CartItemControllerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type cartItemController struct {
	service service.CartItemService
}

func NewCartItemController(service service.CartItemService) CartItemControllerInterface {
	return &cartItemController{service}
}

func (cartItemController *cartItemController) Create(c *gin.Context) {
	customerId := c.MustGet("currentCustomer").(model.Customer).ID

	var input request.CartItemCreateRequest
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	cartItem, err := cartItemController.service.Create(input, int(customerId))
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save cartItem !", http.StatusOK, helper.SUCCESS, response.FormatCart(cartItem))
	c.JSON(http.StatusOK, response)
	return
}

func (cartItemController *cartItemController) Update(c *gin.Context) {
	customerId := c.MustGet("curentCustomer").(model.Customer).ID

	var inputID request.CartItemFindById
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	var inputData request.CartItemUpdateRequest
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	updatedCartItem, err := cartItemController.service.Update(inputID, inputData, int(customerId))
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create order item", http.StatusOK, helper.SUCCESS, response.FormatCart(updatedCartItem))
	c.JSON(http.StatusOK, response)
}

func (cartItemController *cartItemController) FindById(c *gin.Context) {
	var input request.CartItemFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	cartItem, err := cartItemController.service.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("Success get order item !", http.StatusOK, helper.SUCCESS, response.FormatCartItem(cartItem))
	c.JSON(http.StatusOK, response)
}

func (cartItemController *cartItemController) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	cartItems, err := cartItemController.service.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of order items", http.StatusOK, helper.SUCCESS, response.FormatCartItems(cartItems))
	c.JSON(http.StatusOK, response)
}

func (cartItemController *cartItemController) Delete(c *gin.Context) {
	var input request.CartItemFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete cartItem", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	cartItem, err := cartItemController.service.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete cartItem", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete cartItem !", http.StatusOK, helper.SUCCESS, response.FormatCartItem(cartItem))
	c.JSON(http.StatusOK, response)
}
