package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/payment"
	"bwastartup/transaction"
	"bwastartup/user"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectDatabase() (*gorm.DB, error) {
	dsn := "root:End291103#@tcp(mysql_db:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	var db *gorm.DB
	var err error

	// Retry loop untuk memastikan koneksi ke database
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Database connected successfully!")
			return db, nil
		}

		log.Println("Database connection failed. Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	return nil, err
}
func main() {
	db, err := connectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database after multiple attempts:", err)
	}

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService := auth.NewService()
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)

	// user, _ := userService.GetUserByID(3)
	// input := transaction.CreateTransactionInput{
	// 	CampaignID: 1,
	// 	Amount:     10000,
	// 	User:       user,
	// }

	// transactionService.CreateTransaction(input)

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	router.Static("images", "./images")
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadImage)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.GET("campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)

	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.POST("/transactions/notification", (transactionHandler.GetNotification))
	router.Run()

}

// input dari user
// handller , mapping input dari user -> struct input dari user
// Service : melakukan mapping dari stuct input ke struct user ->
// repository
// db
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer ") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrayString := strings.Split(authHeader, " ")
		if len(arrayString) == 2 {
			tokenString = arrayString[1]

		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)

	}

}

//ambil nilai header Authorization : Bearer tokentoken dari GenerateToken
//dari header Authorization , kita ambil nilai tokennya saja
//kita validasi token
//kita ambil user_id
//ambil user dari db berdasarkan user_id lewat servicce
// kita set context isinya user
