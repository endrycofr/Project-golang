package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaigHandler struct {
	service campaign.Service
}

func NewCampaignHandler(s campaign.Service) *campaigHandler {
	return &campaigHandler{service: s}
}

// route api/v1/campaign
func (h *campaigHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.ApiResponse("Error to get campaign", http.StatusBadRequest, "Bad request", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List of campaign", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)

}

// tangkap parameter di handler
// handler ke service
// service yang menentukan repo mana yang dipanggil
//repo : FindAll , FindByUserID
//db
