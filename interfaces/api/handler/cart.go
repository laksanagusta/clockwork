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

type CartHandlerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
	CheckActiveCart(c *gin.Context)
}

type cartHandler struct {
	application application.CartService
}

func NewCartHandler(application application.CartService) CartHandlerInterface {
	return &cartHandler{application}
}

func (cartCN *cartHandler) Create(c *gin.Context) {
	userId := c.MustGet("currentUser").(model.User).ID

	cart, err := cartCN.application.Create(int(userId))
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save cart !", http.StatusOK, helper.SUCCESS, response.FormatCart(cart))
	c.JSON(http.StatusOK, response)
	return
}

func (cartCN *cartHandler) Update(c *gin.Context) {
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

	updatedCart, err := cartCN.application.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create order item", http.StatusOK, helper.SUCCESS, response.FormatCart(updatedCart))
	c.JSON(http.StatusOK, response)
}

func (cartCN *cartHandler) FindById(c *gin.Context) {
	var input request.CartFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	cart, err := cartCN.application.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
		return
	}

	response := helper.APIResponse("Success get order item !", http.StatusOK, helper.SUCCESS, response.FormatCart(cart))
	c.JSON(http.StatusOK, response)
}

func (cartCN *cartHandler) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	carts, err := cartCN.application.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of order items", http.StatusOK, helper.SUCCESS, response.FormatCarts(carts))
	c.JSON(http.StatusOK, response)
}

func (cartCN *cartHandler) Delete(c *gin.Context) {
	var input request.CartFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete cart", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	cart, err := cartCN.application.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete cart", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete cart !", http.StatusOK, helper.SUCCESS, response.FormatCart(cart))
	c.JSON(http.StatusOK, response)
}

func (cartCN *cartHandler) CheckActiveCart(c *gin.Context) {
	userId := c.MustGet("currentUser").(model.User).ID
	cart, err := cartCN.application.CheckActiveCart(int(userId))
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
		return
	}

	response := helper.APIResponse("Success get active cart !", http.StatusOK, helper.SUCCESS, response.FormatCart(cart))
	c.JSON(http.StatusOK, response)
}
