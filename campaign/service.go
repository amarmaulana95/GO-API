package campaign

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignsByID(input GetCampainDetailInput) (Campaign, error) //param input.go
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {

	if userID != 0 {
		// clien get data campaign berdasarkan id user
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	// else ambil smua
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignsByID(input GetCampainDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindbyID(input.ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
