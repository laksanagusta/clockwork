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

type CartItemHandlerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type cartItemHandler struct {
	application application.CartItemService
}

func NewCartItemHandler(application application.CartItemService) CartItemHandlerInterface {
	return &cartItemHandler{application}
}

func (cartItemHandler *cartItemHandler) Create(c *gin.Context) {
	customerId := c.MustGet("currentCustomer").(model.Customer).ID

	var input request.CartItemCreateRequest
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	cartItem, err := cartItemHandler.application.Create(input, int(customerId))
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save cartItem !", http.StatusOK, helper.SUCCESS, response.FormatCart(cartItem))
	c.JSON(http.StatusOK, response)
	return
}

func (cartItemHandler *cartItemHandler) Update(c *gin.Context) {
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

	updatedCartItem, err := cartItemHandler.application.Update(inputID, inputData, int(customerId))
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create order item", http.StatusOK, helper.SUCCESS, response.FormatCart(updatedCartItem))
	c.JSON(http.StatusOK, response)
}

func (cartItemHandler *cartItemHandler) FindById(c *gin.Context) {
	var input request.CartItemFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	cartItem, err := cartItemHandler.application.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("Success get order item !", http.StatusOK, helper.SUCCESS, response.FormatCartItem(cartItem))
	c.JSON(http.StatusOK, response)
}

func (cartItemHandler *cartItemHandler) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	cartItems, err := cartItemHandler.application.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of order items", http.StatusOK, helper.SUCCESS, response.FormatCartItems(cartItems))
	c.JSON(http.StatusOK, response)
}

func (cartItemHandler *cartItemHandler) Delete(c *gin.Context) {
	var input request.CartItemFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete cartItem", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	cartItem, err := cartItemHandler.application.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete cartItem", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete cartItem !", http.StatusOK, helper.SUCCESS, response.FormatCartItem(cartItem))
	c.JSON(http.StatusOK, response)
}
