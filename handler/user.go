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

func (h *userHandler) Login(c *gin.Context) {
// user memasukkan input (email, password)
//input ditangkap handler
//mapping dari input user ke input struct
// input struct passing service
// di service mencari dg bantuan repository user dengan email x
// mencocokan password 

var input user.LoginInput
err := c.ShouldBind(&input)
if err != nil {
	errors := helper.FormatValidationError(err)
	errorMassage := gin.H{"error": errors}
	response := helper.ApiResponse("Login Failed ", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loggedinuser, err := h.userService.Login(input)
	if err != nil{
		errorMassage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Login Failed ", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter :=user.FormatterUser(loggedinuser,"tokentokentoken")
	response := helper.ApiResponse("Successfully loggedin", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}