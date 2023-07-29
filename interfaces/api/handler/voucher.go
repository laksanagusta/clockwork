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
	ApplyVoucher(c *gin.Context)
}

type voucherHandler struct {
	application   application.VoucherService
	globalHelper  helper.GlobalHelper
	voucherHelper helper.VoucherHelper
}

func NewVoucherHandler(application application.VoucherService, globalHelper helper.GlobalHelper) VoucherHandlerInterface {
	voucherHelper := helper.NewVoucherHelper()
	return &voucherHandler{application,
		globalHelper,
		voucherHelper,
	}
}

func (voucherHandler *voucherHandler) Create(c *gin.Context) {
	var input request.VoucherCreateInput
	err := c.ShouldBind(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	err = voucherHandler.voucherHelper.ValidateTimeValid(input.StartTime, input.EndTime)
	if err != nil {
		helper.ErrorResponse(err, c, helper.VALIDATION_ERROR_MESSAGE)
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

	err = voucherHandler.voucherHelper.ValidateTimeValid(inputData.StartTime, inputData.EndTime)
	if err != nil {
		helper.ErrorResponse(err, c, helper.VALIDATION_ERROR_MESSAGE)
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

func (voucherHandler *voucherHandler) ApplyVoucher(c *gin.Context) {
	var input request.VoucherApply
	err := c.ShouldBindJSON(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	cart, err := voucherHandler.application.ApplyVoucher(input)
	if err != nil {
		helper.ErrorResponse(err, c, "Failed to apply voucher")
		return
	}

	response := helper.APIResponse("Success apply voucher !", http.StatusOK, helper.SUCCESS, response.FormatCart(cart))
	c.JSON(http.StatusOK, response)
}
