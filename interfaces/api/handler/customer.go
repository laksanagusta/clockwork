package handler

import (
	"clockwork-server/application"
	"clockwork-server/domain/model"
	"clockwork-server/helper"
	"clockwork-server/interfaces/api/auth"
	"clockwork-server/interfaces/api/request"
	"clockwork-server/interfaces/api/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	customerService application.CustomerService
	authService     auth.Auth
}

func NewCustomerHandler(customerService application.CustomerService, authService auth.Auth) *customerHandler {
	return &customerHandler{customerService, authService}
}

func (handler *customerHandler) RegisterCustomer(c *gin.Context) {
	var input request.RegisterCustomerRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	newCustomer, err := handler.customerService.RegisterCustomer(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.REGISTER_FAILED_MESSAGE)
		return
	}

	token, err := handler.authService.GenerateToken(int(newCustomer.ID), newCustomer.Email, "customer")
	if err != nil {
		helper.ErrorResponse(err, c, helper.REGISTER_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", response.FormatCustomer(newCustomer, token))
	c.JSON(http.StatusOK, response)
	return
}

func (handler *customerHandler) Login(c *gin.Context) {
	var input request.LoginRequest
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	loggedinCustomer, err := handler.customerService.Login(input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.LOGIN_FAILED_MESSAGE)
	}

	token, err := handler.authService.GenerateToken(int(loggedinCustomer.ID), loggedinCustomer.Email, "customer")
	if err != nil {
		helper.ErrorValidation(err, c, helper.LOGIN_FAILED_MESSAGE+", Error generating token")
	}

	response := helper.APIResponse("Logged in", http.StatusOK, "success", response.FormatCustomer(loggedinCustomer, token))

	c.JSON(http.StatusOK, response)
}

func (handler *customerHandler) CustomerDetails(c *gin.Context) {
	currentCustomer := c.MustGet("currentCustomer").(model.Customer)
	customer_id := currentCustomer.ID

	customer, err := handler.customerService.CustomerDetails(customer_id)
	if err != nil {
		helper.ErrorValidation(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	c.JSON(http.StatusOK, customer)
}

func (handler *customerHandler) CustomerFindAll(c *gin.Context) {
	customerAll, err := handler.customerService.CustomerFindAll()
	if err != nil {
		helper.ErrorValidation(err, c, helper.SOMETHING_WENT_WRONG_MESSAGE)
	}

	response := helper.APIResponse("List of customers", http.StatusOK, "success", response.FormatCustomers(customerAll))
	c.JSON(http.StatusOK, response)
}

func (handler *customerHandler) UpdateCustomer(c *gin.Context) {
	var inputID request.CustomerFindById
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	var inputData request.UpdateCustomerRequest
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	updatedCustomer, err := handler.customerService.UpdateCustomer(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
	}

	response := helper.APIResponse("Success create customer", http.StatusOK, "success", response.FormatCustomer(updatedCustomer, "token"))
	c.JSON(http.StatusOK, response)
}

func (handler *customerHandler) DeleteCustomer(c *gin.Context) {
	var inputID request.CustomerFindById
	err := c.ShouldBindUri(&inputID)

	deletedCustomer, err := handler.customerService.DeleteCustomer(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete customer", http.StatusOK, "success", deletedCustomer)
	c.JSON(http.StatusOK, response)

}
