package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	
	
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	router.Run()

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "fiki"
	// userInput.Email = "fikians@gmail.com"
	// userInput.Occupation = "model"
	// userInput.Password = "lempuyangansambelijo"
	// userService.RegisterUser(userInput)
	
}

// input dari user
// handller , mapping input dari user -> struct input dari user
// Service : melakukan mapping dari stuct input ke struct user ->
//repository
//db

