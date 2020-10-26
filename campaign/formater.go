package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	userID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	//buat object
	CampaignFormatter := CampaignFormatter{}
	CampaignFormatter.ID = campaign.ID
	CampaignFormatter.userID = campaign.userID
	CampaignFormatter.Name = campaign.Name
	CampaignFormatter.ShortDescription = campaign.ShortDescription
	CampaignFormatter.GoalAmount = campaign.GoalAmount
	CampaignFormatter.CurrentAmount = campaign.CurrentAmount
	CampaignFormatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		CampaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}
	return CampaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	//array campaign formater
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}
	return campaignsFormatter
}
