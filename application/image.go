package application

import (
	"clockwork-server/domain/model"
	"clockwork-server/domain/repository"
	"clockwork-server/interfaces/api/request"
	"errors"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
)

type ImageService interface {
	Create(request request.ImageCreateRequest, file *multipart.FileHeader) (model.Image, error)
	Remove(request request.ImageRemoveRequest) (model.Image, error)
	Update(request request.ImageUpdateRequest, params request.ImageFindByIdRequest) (model.Image, error)
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

func (s *imageService) Update(request request.ImageUpdateRequest, params request.ImageFindByIdRequest) (model.Image, error) {
	image, err := s.imageRepo.FindById(params.ID)
	if err != nil {
		return image, err
	}

	image.IsPrimary = request.IsPrimary

	if request.IsPrimary == true {
		err = s.imageRepo.UpdateIsPrimaryFalse(int(image.ProductID))
		if err != nil {
			return image, err
		}
	}

	updateImage, err := s.imageRepo.Update(image)
	if err != nil {
		return image, err
	}

	return updateImage, err
}

func (s *imageService) Remove(request request.ImageRemoveRequest) (model.Image, error) {
	image, err := s.imageRepo.FindById(request.ID)
	if err != nil {
		return image, err
	}

	err = os.Remove(image.Url)
	if err != nil {
		return image, err
	}

	_, err = s.imageRepo.Remove(request.ID)
	if err != nil {
		return image, err
	}

	return image, nil
}
