package handler

import (
	"gomar/campaign"
	"gomar/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// tangkap param di handler
// handler ke service
// service menentukan repository mana yg akan di call
// repo : GetAll GetUserbyID
// db

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

//api/v1/campaigns
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id")) //ubah ke int dngn strconv.atoi

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error get data campaigns ", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List data campaigns ", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusBadRequest, response)
	return
}

func (h *campaignHandler) GetCampaign(c *gin.Context) { //detail campaign
	// api/v1/campaigns/ID
	// handler : maping id yg di url ke struct input => service, call formatter
	// service : input struct input => tangkap id di url, panggil repo
	// butuh repository: u get campaign id

	var input campaign.GetCampainDetailInput
	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponse("Failed get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaign, err := h.service.GetCampaignsByID(input)
	if err != nil {
		response := helper.APIResponse("Failed get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("campaign detail", http.StatusOK, "success", campaign)
	c.JSON(http.StatusOK, response)
}
