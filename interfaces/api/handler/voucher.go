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

type VoucherHandlerInterface interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	Delete(c *gin.Context)
}

type voucherHandler struct {
	application application.VoucherService
}

func NewVoucherHandler(application application.VoucherService) VoucherHandlerInterface {
	return &voucherHandler{application}
}

func (voucherHandler *voucherHandler) Create(c *gin.Context) {
	var input request.VoucherCreateInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	voucher, err := voucherHandler.application.Create(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success save voucher !", http.StatusOK, helper.SUCCESS, response.FormatVoucher(voucher))
	c.JSON(http.StatusOK, response)
	return
}

func (voucherHandler *voucherHandler) Update(c *gin.Context) {
	var inputID request.VoucherFindById
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	var inputData request.VoucherUpdateInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	updatedVoucher, err := voucherHandler.application.Update(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create voucher", http.StatusOK, helper.SUCCESS, response.FormatVoucher(updatedVoucher))
	c.JSON(http.StatusOK, response)
}

func (voucherHandler *voucherHandler) FindById(c *gin.Context) {
	var input request.VoucherFindById

	err := c.ShouldBindUri(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	voucher, err := voucherHandler.application.FindById(input.ID)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("Success get voucher !", http.StatusOK, helper.SUCCESS, response.FormatVoucher(voucher))
	c.JSON(http.StatusOK, response)
}

func (voucherHandler *voucherHandler) FindAll(c *gin.Context) {
	q := c.Request.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	pageSize, _ := strconv.Atoi(q.Get("limit"))
	s := q.Get("q")
	vouchers, err := voucherHandler.application.FindAll(page, pageSize, s)
	if err != nil {
		helper.ErrorResponse(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	response := helper.APIResponse("List of vouchers", http.StatusOK, helper.SUCCESS, response.FormatVouchers(vouchers))
	c.JSON(http.StatusOK, response)
}

func (voucherHandler *voucherHandler) Delete(c *gin.Context) {
	var input request.VoucherFindById
	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("Failed to delete voucher", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	voucher, err := voucherHandler.application.Delete(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete voucher", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete voucher !", http.StatusOK, helper.SUCCESS, response.FormatVoucher(voucher))
	c.JSON(http.StatusOK, response)
}
