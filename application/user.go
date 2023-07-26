package application

import (
	"clockwork-server/domain/model"
	"clockwork-server/domain/repository"
	"clockwork-server/interfaces/api/request"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(input request.RegisterUserRequest) (model.User, error)
	Login(input request.LoginRequest) (model.User, error)
	UserDetails(id uint64) (model.User, error)
	UserFindAll() ([]model.User, error)
	UpdateUser(inputID request.GetUserDetailRequest, inputData request.UpdateUserRequest) (model.User, error)
	DeleteUser(inputID request.GetUserDetailRequest) (string, error)
	CreateUserBulk(workerIndex int, counter int, jobs []interface{}) (model.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository}
}

func (s *userService) RegisterUser(input request.RegisterUserRequest) (model.User, error) {
	user := model.User{}
	user.Username = strings.ToLower(input.Username)
	user.Name = input.Name
	user.Email = strings.ToLower(input.Email)
	user.Occupation = input.Occupation
	user.PhoneNumber = input.PhoneNumber
	user.Address = input.Address

	isEmailExist, err := s.repository.FindByEmail(user.Email)
	if err != nil {
		return user, err
	}

	if isEmailExist.ID != 0 {
		return user, errors.New("Register failed, Email already exist")
	}

	isUsernameExist, err := s.repository.FindByUsername(user.Username)
	if err != nil {
		return user, err
	}

	if isUsernameExist.ID != 0 {
		return user, errors.New("Register failed, username already been taken")
	}

	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(PasswordHash)

	user.Role = "USER"
	if input.Role != "" {
		user.Role = strings.ToUpper(input.Role)
	}

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *userService) Login(input request.LoginRequest) (model.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) UserDetails(id uint64) (model.User, error) {
	user_id := id
	user, err := s.repository.FindById(user_id)
	if user.ID == 0 {
		return user, errors.New("User not found")
	}
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) UserFindAll() ([]model.User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *userService) UpdateUser(inputID request.GetUserDetailRequest, inputData request.UpdateUserRequest) (model.User, error) {
	user, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return user, err
	}

	user.Username = strings.ToLower(inputData.Username)
	user.Name = inputData.Name
	user.Email = strings.ToLower(inputData.Email)
	user.Occupation = inputData.Occupation
	user.Role = "USER"

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil

}

func (s *userService) DeleteUser(inputID request.GetUserDetailRequest) (string, error) {
	updatedUser, err := s.repository.Delete(inputID.ID)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *userService) UpdateUserRole(inputID request.GetUserDetailRequest, inputData request.UpdateUserRequest) (model.User, error) {
	user, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return user, err
	}

	user.Name = inputData.Name
	user.Email = inputData.Email
	user.Occupation = inputData.Occupation
	user.Role = "ADMIN"

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil

}

func (s *userService) CreateUserBulk(workerIndex int, counter int, jobs []interface{}) (model.User, error) {
	user := model.User{}
	user.Name = fmt.Sprintf("%v", jobs[0])

	newUsers, err := s.repository.Save(user)
	if err != nil {
		return newUsers, err
	}

	return newUsers, nil
}
