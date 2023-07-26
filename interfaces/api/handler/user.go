package handler

import (
	"clockwork-server/application"
	"clockwork-server/domain/model"
	"clockwork-server/helper"
	"clockwork-server/interfaces/api/auth"
	"clockwork-server/interfaces/api/request"
	"clockwork-server/interfaces/api/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService application.UserService
	authService auth.Auth
}

func NewUserHandler(userService application.UserService, authService auth.Auth) *userHandler {
	return &userHandler{userService, authService}
}

func (handler *userHandler) RegisterUser(c *gin.Context) {
	var input request.RegisterUserRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
		return
	}

	newUser, err := handler.userService.RegisterUser(input)
	if err != nil {
		helper.ErrorResponse(err, c, helper.REGISTER_FAILED_MESSAGE)
		return
	}

	token, err := handler.authService.GenerateToken(int(newUser.ID), newUser.Email, "user")
	if err != nil {
		helper.ErrorResponse(err, c, helper.REGISTER_FAILED_MESSAGE)
		return
	}

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", response.FormatUser(newUser, token))
	c.JSON(http.StatusOK, response)
	return
}

func (handler *userHandler) Login(c *gin.Context) {
	var input request.LoginRequest
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	loggedinUser, err := handler.userService.Login(input)
	if err != nil {
		helper.ErrorValidation(err, c, helper.LOGIN_FAILED_MESSAGE)
	}

	token, err := handler.authService.GenerateToken(int(loggedinUser.ID), loggedinUser.Email, "user")
	if err != nil {
		helper.ErrorValidation(err, c, helper.LOGIN_FAILED_MESSAGE+", Error generating token")
	}

	response := helper.APIResponse("Logged in", http.StatusOK, "success", response.FormatUser(loggedinUser, token))

	c.JSON(http.StatusOK, response)
}

func (handler *userHandler) UserDetails(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(model.User)
	user_id := currentUser.ID
	id := c.Param("id")

	user_id, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helper.ErrorValidation(err, c, helper.SOMETHING_WENT_WRONG_MESSAGE)
	}

	user, err := handler.userService.UserDetails(user_id)
	if err != nil {
		helper.ErrorValidation(err, c, helper.FAILED_GET_DATA_MESSAGE)
	}

	c.JSON(http.StatusOK, user)
}

func (handler *userHandler) UserFindAll(c *gin.Context) {
	userAll, err := handler.userService.UserFindAll()
	if err != nil {
		helper.ErrorValidation(err, c, helper.SOMETHING_WENT_WRONG_MESSAGE)
	}

	response := helper.APIResponse("List of users", http.StatusOK, "success", response.FormatUsers(userAll))
	c.JSON(http.StatusOK, response)
}

func (handler *userHandler) UpdateUser(c *gin.Context) {
	var inputID request.GetUserDetailRequest
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	var inputData request.UpdateUserRequest
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		helper.ErrorValidation(err, c, helper.VALIDATION_ERROR_MESSAGE)
	}

	updatedUser, err := handler.userService.UpdateUser(inputID, inputData)
	if err != nil {
		helper.ErrorResponse(err, c, helper.UPDATE_FAILED_MESSAGE)
	}

	response := helper.APIResponse("Success create user", http.StatusOK, "success", response.FormatUser(updatedUser, "token"))
	c.JSON(http.StatusOK, response)
}

func (handler *userHandler) DeleteUser(c *gin.Context) {
	var inputID request.GetUserDetailRequest
	err := c.ShouldBindUri(&inputID)

	deletedUser, err := handler.userService.DeleteUser(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete user", http.StatusOK, "success", deletedUser)
	c.JSON(http.StatusOK, response)

}
