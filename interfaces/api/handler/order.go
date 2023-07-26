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

const SAVE_ORDER_FAILED = "Failed to save order"
const GET_ORDER_FAILED = "Failed to get order"

type OrderHandlerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type orderHandler struct {
	application application.OrderService
}

func NewOrderHandler(application application.OrderService) OrderHandlerInterface {
	return &orderHandler{application}
}

func (_orderHandler *orderHandler) Create(c *gin.Context) {
	var input request.OrderCreateRequest
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse(SAVE_ORDER_FAILED, http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	order, err := _orderHandler.application.Create(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse(SAVE_ORDER_FAILED, http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success save order !", http.StatusOK, helper.SUCCESS, response.FormatOrder(order))
	c.JSON(http.StatusOK, response)
}

func (_orderHandler *orderHandler) Update(c *gin.Context) {
	var inputID request.OrderFindById
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse(SAVE_ORDER_FAILED, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData request.OrderUpdateRequest
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse(SAVE_ORDER_FAILED, http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedOrder, err := _orderHandler.application.Update(inputID, inputData)
	if err != nil {
		response := helper.APIResponse(SAVE_ORDER_FAILED, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create order", http.StatusOK, helper.SUCCESS, response.FormatOrder(updatedOrder))
	c.JSON(http.StatusOK, response)
}

func (_orderHandler *orderHandler) FindById(c *gin.Context) {
	var input request.OrderFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse(GET_ORDER_FAILED, http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	order, err := _orderHandler.application.FindById(input.ID)
	if err != nil {
		response := helper.APIResponse(GET_ORDER_FAILED, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get order !", http.StatusOK, helper.SUCCESS, response.FormatOrder(order))
	c.JSON(http.StatusOK, response)
}

func (_orderHandler *orderHandler) FindByCode(c *gin.Context) {
	var input request.OrderFindByCode
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse(GET_ORDER_FAILED, http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	orderSingle, err := _orderHandler.application.FindByCode(input.Code)
	if err != nil {
		response := helper.APIResponse(GET_ORDER_FAILED, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get order !", http.StatusOK, helper.SUCCESS, response.FormatOrder(orderSingle))
	c.JSON(http.StatusOK, response)
}

func (_orderHandler *orderHandler) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	s := q.Get("q")
	orders, err := _orderHandler.application.FindAll(page, pageSize, s)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse(GET_ORDER_FAILED, http.StatusOK, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of orders", http.StatusOK, helper.SUCCESS, response.FormatOrders(orders))
	c.JSON(http.StatusOK, response)
}

func (_orderHandler *orderHandler) Delete(c *gin.Context) {
	var input request.OrderFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete order", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	order, err := _orderHandler.application.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete order !", http.StatusOK, helper.SUCCESS, response.FormatOrder(order))
	c.JSON(http.StatusOK, response)
}
