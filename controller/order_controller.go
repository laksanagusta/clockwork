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

const SAVE_ORDER_FAILED = "Failed to save order"
const GET_ORDER_FAILED = "Failed to get order"

type OrderControllerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type orderController struct {
	service service.OrderService
}

func NewOrderController(service service.OrderService) OrderControllerInterface {
	return &orderController{service}
}

func (_orderController *orderController) Create(c *gin.Context) {
	var input request.OrderCreateRequest
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse(SAVE_ORDER_FAILED, http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	order, err := _orderController.service.Create(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse(SAVE_ORDER_FAILED, http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success save order !", http.StatusOK, helper.SUCCESS, response.FormatOrder(order))
	c.JSON(http.StatusOK, response)
}

func (_orderController *orderController) Update(c *gin.Context) {
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

	updatedOrder, err := _orderController.service.Update(inputID, inputData)
	if err != nil {
		response := helper.APIResponse(SAVE_ORDER_FAILED, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create order", http.StatusOK, helper.SUCCESS, response.FormatOrder(updatedOrder))
	c.JSON(http.StatusOK, response)
}

func (_orderController *orderController) FindById(c *gin.Context) {
	var input request.OrderFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse(GET_ORDER_FAILED, http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	order, err := _orderController.service.FindById(input.ID)
	if err != nil {
		response := helper.APIResponse(GET_ORDER_FAILED, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get order !", http.StatusOK, helper.SUCCESS, response.FormatOrder(order))
	c.JSON(http.StatusOK, response)
}

func (_orderController *orderController) FindByCode(c *gin.Context) {
	var input request.OrderFindByCode
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse(GET_ORDER_FAILED, http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	orderSingle, err := _orderController.service.FindByCode(input.Code)
	if err != nil {
		response := helper.APIResponse(GET_ORDER_FAILED, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success get order !", http.StatusOK, helper.SUCCESS, response.FormatOrder(orderSingle))
	c.JSON(http.StatusOK, response)
}

func (_orderController *orderController) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	s := q.Get("q")
	orders, err := _orderController.service.FindAll(page, pageSize, s)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse(GET_ORDER_FAILED, http.StatusOK, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of orders", http.StatusOK, helper.SUCCESS, response.FormatOrders(orders))
	c.JSON(http.StatusOK, response)
}

func (_orderController *orderController) Delete(c *gin.Context) {
	var input request.OrderFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete order", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	order, err := _orderController.service.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete order", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete order !", http.StatusOK, helper.SUCCESS, response.FormatOrder(order))
	c.JSON(http.StatusOK, response)
}
