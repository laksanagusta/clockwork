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

type OrderItemControllerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type orderItemController struct {
	service service.OrderItemService
}

func NewOrderItemController(service service.OrderItemService) OrderItemControllerInterface {
	return &orderItemController{service}
}

func (orderItemController *orderItemController) Create(c *gin.Context) {
	customerId := c.MustGet("curentCustomer").(model.Customer).ID

	var input request.OrderItemCreateInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	orderItem, err := orderItemController.service.Create(input, int(customerId))
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save orderItem !", http.StatusOK, helper.SUCCESS, response.FormatOrder(orderItem))
	c.JSON(http.StatusOK, response)
	return
}

func (orderItemController *orderItemController) Update(c *gin.Context) {
	var inputID request.OrderItemFindById
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	var inputData request.OrderItemUpdateInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	updatedOrderItem, err := orderItemController.service.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create order item", http.StatusOK, helper.SUCCESS, response.FormatOrder(updatedOrderItem))
	c.JSON(http.StatusOK, response)
}

func (orderItemController *orderItemController) FindById(c *gin.Context) {
	var input request.OrderItemFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	orderItem, err := orderItemController.service.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("Success get order item !", http.StatusOK, helper.SUCCESS, response.FormatOrderItem(orderItem))
	c.JSON(http.StatusOK, response)
}

func (orderItemController *orderItemController) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	orderItems, err := orderItemController.service.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of order items", http.StatusOK, helper.SUCCESS, response.FormatOrderItems(orderItems))
	c.JSON(http.StatusOK, response)
}

func (orderItemController *orderItemController) Delete(c *gin.Context) {
	var input request.OrderItemFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete orderItem", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	orderItem, err := orderItemController.service.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete orderItem", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete orderItem !", http.StatusOK, helper.SUCCESS, response.FormatOrderItem(orderItem))
	c.JSON(http.StatusOK, response)
}
