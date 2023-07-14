package controller

import (
	"clockwork-server/helper"
	"clockwork-server/request"
	"clockwork-server/response"
	"clockwork-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageController interface {
	Create(c *gin.Context)
}

type imageController struct {
	imageService service.ImageService
}

func NewImageController(imageService service.ImageService) ImageController {
	return &imageController{imageService}
}

func (ic *imageController) Create(c *gin.Context) {
	var input request.ImageCreateRequest
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	file, err := c.FormFile("file")
	if err != nil {
		helper.ErrorResponse(err, c, "No file is received")
	}

	image, err := ic.imageService.Create(input, file)
	if err != nil {
		helper.ErrorResponse(err, c, helper.SAVE_FAILED_MESSAGE)
	}

	response := helper.APIResponse("Success create image", http.StatusOK, helper.SUCCESS, response.FormatImage(image))
	c.JSON(http.StatusOK, response)
}
