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
