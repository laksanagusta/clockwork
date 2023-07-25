package service

import (
	"clockwork-server/model"
	"clockwork-server/repository"
	"clockwork-server/request"
	"errors"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
)

type ImageService interface {
	Create(request request.ImageCreateRequest, file *multipart.FileHeader) (model.Image, error)
	Remove(request request.ImageRemoveRequest) (int8, error)
}

type imageService struct {
	imageRepo repository.ImageRepository
}

func NewImageService(imageRepo repository.ImageRepository) ImageService {
	return &imageService{imageRepo}
}

func (s *imageService) Create(request request.ImageCreateRequest, file *multipart.FileHeader) (model.Image, error) {
	image := model.Image{}

	if file.Size > 2000000 {
		return image, errors.New("File exceeds 2mb")
	}

	images, err := s.imageRepo.GetImagesByProductId(request.ProductID)
	if err != nil {
		return image, err
	}

	if len(images) == 4 {
		return image, errors.New("Can't add more image, image already 4, please remove some")
	}

	if len(images) > 0 && request.IsPrimary == true {
		err = s.imageRepo.UpdateIsPrimaryFalse(request.ProductID)
		if err != nil {
			return image, err
		}
	}

	// extension := filepath.Ext(file.Filename)
	filename := file.Filename

	path := fmt.Sprintf("images/%d-%s", request.ProductID, filename)

	c := gin.Context{}
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		return image, err
	}

	image.ProductID = uint(request.ProductID)
	image.Url = path
	image.IsPrimary = request.IsPrimary

	newImage, err := s.imageRepo.Create(image)
	if err != nil {
		return newImage, err
	}

	return newImage, nil
}

func (s *imageService) Remove(request request.ImageRemoveRequest) (int8, error) {
	image, err := s.imageRepo.FindById(request.ID)
	if err != nil {
		return request.ID, err
	}

	err = os.Remove(image.Url)
	if err != nil {
		return request.ID, err
	}

	_, err = s.imageRepo.Remove(request.ID)
	if err != nil {
		return request.ID, err
	}

	return request.ID, nil
}
