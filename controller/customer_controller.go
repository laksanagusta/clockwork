package controller

import (
	"clockwork-server/auth"
	"clockwork-server/helper"
	"clockwork-server/model"
	"clockwork-server/request"
	"clockwork-server/response"
	"clockwork-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type customerController struct {
	customerService service.CustomerService
	authService     auth.Auth
}

func NewCustomerController(customerService service.CustomerService, authService auth.Auth) *customerController {
	return &customerController{customerService, authService}
}

func (controller *customerController) RegisterCustomer(c *gin.Context) {
	var input request.RegisterCustomerRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	newCustomer, err := controller.customerService.RegisterCustomer(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.REGISTER_FAILED_MESSAGE)
		return
	}

	token, err := controller.authService.GenerateToken(uint64(newCustomer.ID))
	if err != nil {
		helper.ErrorResponse(err, c, helper.REGISTER_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", response.FormatCustomer(newCustomer, token))
	c.JSON(http.StatusOK, response)
	return
}

func (controller *customerController) Login(c *gin.Context) {
	var input request.LoginRequest
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	loggedinCustomer, err := controller.customerService.Login(input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.LOGIN_FAILED_MESSAGE)
	}

	token, err := controller.authService.GenerateToken(uint64(loggedinCustomer.ID))
	if err != nil {
		helper.ErrorValidation(err, c, helper.LOGIN_FAILED_MESSAGE+", Error generating token")
	}

	response := helper.APIResponse("Logged in", http.StatusOK, "success", response.FormatCustomer(loggedinCustomer, token))

	c.JSON(http.StatusOK, response)
}

func (controller *customerController) CustomerDetails(c *gin.Context) {
	currentCustomer := c.MustGet("currentCustomer").(model.Customer)
	customer_id := currentCustomer.ID

	customer, err := controller.customerService.CustomerDetails(customer_id)
	if err != nil {
		helper.ErrorValidation(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	c.JSON(http.StatusOK, customer)
}

func (controller *customerController) CustomerFindAll(c *gin.Context) {
	customerAll, err := controller.customerService.CustomerFindAll()
	if err != nil {
		helper.ErrorValidation(err, c, helper.SOMETHING_WENT_WRONG_MESSAGE)
	}

	response := helper.APIResponse("List of customers", http.StatusOK, "success", response.FormatCustomers(customerAll))
	c.JSON(http.StatusOK, response)
}

func (controller *customerController) UpdateCustomer(c *gin.Context) {
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

	updatedCustomer, err := controller.customerService.UpdateCustomer(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
	}

	response := helper.APIResponse("Success create customer", http.StatusOK, "success", response.FormatCustomer(updatedCustomer, "token"))
	c.JSON(http.StatusOK, response)
}

func (controller *customerController) DeleteCustomer(c *gin.Context) {
	var inputID request.CustomerFindById
	err := c.ShouldBindUri(&inputID)

	deletedCustomer, err := controller.customerService.DeleteCustomer(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete customer", http.StatusOK, "success", deletedCustomer)
	c.JSON(http.StatusOK, response)

}
