package service

import (
	"clockwork-server/model"
	"clockwork-server/repository"
	"clockwork-server/request"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type ImageService interface {
	Create(request request.ImageCreateRequest, file *multipart.FileHeader) (model.Image, error)
	Remove(request request.ImageRemoveRequest) (int8, error)
}

type imageService struct {
	imageRepo repository.ImageRepository
	c         *gin.Context
}

func NewImageService(imageRepo repository.ImageRepository, c *gin.Context) ImageService {
	return &imageService{imageRepo, c}
}

func (is *imageService) Create(request request.ImageCreateRequest, file *multipart.FileHeader) (model.Image, error) {
	image := model.Image{}
	image.Url = request.Url

	extension := filepath.Ext(file.Filename)
	filename := request.ProductID + extension

	err := is.c.SaveUploadedFile(file, "/image/product/"+filename)
	if err != nil {
		return image, err
	}

	newImage, err := is.imageRepo.Create(image)
	if err != nil {
		return newImage, err
	}

	return newImage, nil
}

func (is *imageService) Remove(request request.ImageRemoveRequest) (int8, error) {
	newImage, err := is.imageRepo.Remove(request.ID)
	if err != nil {
		return newImage, err
	}

	return newImage, nil
}
