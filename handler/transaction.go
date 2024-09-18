package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// parameter di uri
// tangkap parameter mapping ke input struct
// panggil service , input struct sebagai parameter
// service ,  harus mempunyai CampaignID untuk bisa memanggil repo
// repo mencari data transaction suatu campaign

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service: service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get transactions of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List Campaing's transactions  ", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)

}

// GetUserTransatios (pada saat user login )
// Handler
//ambil nilai user dr jwt (authMiddleware)
// service
// repo => mengambil data transaction (di data itu harus ada preload campaign )

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	transactions, err := h.service.GetTransactionsByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get user's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List User's transactions", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}
