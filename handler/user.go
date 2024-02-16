package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}

}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMassage := gin.H{"error": errors}
		
		response := helper.ApiResponse("Failed register new account", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.ApiResponse("Failed register new account", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}


	// Implement token generation if needed
	// token, err := h.jwtService.GenerateToken()

	formatter := user.FormatterUser(newUser, "tokentokentokentokentoken")
	response := helper.ApiResponse("Success register new account", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
