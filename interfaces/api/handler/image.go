package handler

import (
	"clockwork-server/application"
	"clockwork-server/helper"
	"clockwork-server/interfaces/api/request"
	"clockwork-server/interfaces/api/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageHandler interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

type imageHandler struct {
	imageService application.ImageService
}

func NewImageHandler(imageService application.ImageService) ImageHandler {
	return &imageHandler{imageService}
}

func (ic *imageHandler) Create(c *gin.Context) {
	var input request.ImageCreateRequest

	err := c.ShouldBind(&input)

	if err != nil {
		fmt.Println(err)
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		helper.ErrorResponse(err, c, "No file is received")
		return
	}

	image, err := ic.imageService.Create(input, file)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create image", http.StatusOK, helper.SUCCESS, response.FormatImage(image))
	c.JSON(http.StatusOK, response)
}

func (ic *imageHandler) Update(c *gin.Context) {
	var params request.ImageFindByIdRequest
	err := c.ShouldBindUri(&params)
	if err != nil {
		response := helper.APIResponse(SAVE_ORDER_FAILED, http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input request.ImageUpdateRequest
	err = c.ShouldBindJSON(&input)
	if err != nil {
		fmt.Println(err)
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	image, err := ic.imageService.Update(input, params)
	if err != nil {
		helper.ErrorResponse(err, c, helper.DELETE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success update image", http.StatusOK, helper.SUCCESS, response.FormatImage(image))
	c.JSON(http.StatusOK, response)
}

func (ic *imageHandler) Delete(c *gin.Context) {
	var input request.ImageRemoveRequest
	err := c.ShouldBindUri(&input)

	if err != nil {
		fmt.Println(err)
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	image, err := ic.imageService.Remove(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.DELETE_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Success create image", http.StatusOK, helper.SUCCESS, response.FormatImage(image))
	c.JSON(http.StatusOK, response)
}
